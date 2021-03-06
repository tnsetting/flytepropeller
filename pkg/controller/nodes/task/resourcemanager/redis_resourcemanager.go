package resourcemanager

import (
	"context"
	"strings"
	"sync"
	"time"

	pluginCore "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
	rmConfig "github.com/lyft/flytepropeller/pkg/controller/nodes/task/resourcemanager/config"
	"github.com/lyft/flytestdlib/logger"
	"github.com/lyft/flytestdlib/promutils"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apimachinery/pkg/util/wait"
)

// This is the key that will point to the Redis Set.
// https://redis.io/commands#set
const RedisSetKeyPrefix = "redisresourcemanager"

type RedisResourceManagerBuilder struct {
	client                      RedisClient
	MetricsScope                promutils.Scope
	namespacedResourcesQuotaMap map[pluginCore.ResourceNamespace]int
}

func (r *RedisResourceManagerBuilder) GetID() string {
	return RedisSetKeyPrefix
}

func (r *RedisResourceManagerBuilder) GetResourceRegistrar(namespacePrefix pluginCore.ResourceNamespace) pluginCore.ResourceRegistrar {
	return ResourceRegistrarProxy{
		ResourceRegistrar:       r,
		ResourceNamespacePrefix: namespacePrefix,
	}
}

func (r *RedisResourceManagerBuilder) RegisterResourceQuota(ctx context.Context, namespace pluginCore.ResourceNamespace, quota int) error {
	if r.client == nil {
		err := errors.Errorf("Redis client does not exist.")
		return err
	}

	config := rmConfig.GetConfig()
	if quota <= 0 || quota > config.ResourceMaxQuota {
		err := errors.Errorf("Invalid request for resource quota (<= 0 || > %v): [%v]", config.ResourceMaxQuota, quota)
		return err
	}

	// Checking if the namespace already exists
	// We use linear search here because this function is only called a few times
	if _, ok := r.namespacedResourcesQuotaMap[namespace]; ok {
		return errors.Errorf("Resource namespace already exists [%v]", namespace)
	}

	// Add this registration to the list
	r.namespacedResourcesQuotaMap[namespace] = quota
	logger.Infof(ctx, "Registering resource quota for Namespace [%v]. Quota [%v]", namespace, quota)
	return nil
}

func getValidMetricScopeName(name string) string {
	return strings.Replace(name, "-", "_", -1)
}

func (r *RedisResourceManagerBuilder) BuildResourceManager(ctx context.Context) (pluginCore.ResourceManager, error) {
	if r.client == nil || r.MetricsScope == nil || r.namespacedResourcesQuotaMap == nil {
		return nil, errors.Errorf("Failed to build a redis resource manager. Missing key property(s)")
	}
	logger.Infof(ctx, "Start building a resource manager")
	rm := &RedisResourceManager{
		client:                 r.client,
		MetricsScope:           r.MetricsScope,
		namespacedResourcesMap: map[pluginCore.ResourceNamespace]*Resource{},
	}

	logger.Infof(ctx, "Building a resource manager: creating metrics and namespacedResourcesMap")
	// building the resources and insert them into the resource manager
	for namespace, quota := range r.namespacedResourcesQuotaMap {
		// `namespace` is always prefixed with the RedisSetKeyPrefix and the plugin ID. Each plugin can then affix additional sub-namespaces to it to create different resource pools.
		// For example, hive qubole plugin's namespaces contain plugin ID and qubole cluster (e.g., "redisresourcemanager:qubole-hive-executor:default-cluster").
		metrics := NewRedisResourceManagerMetrics(r.MetricsScope.NewSubScope(getValidMetricScopeName(string(namespace))))

		rm.namespacedResourcesMap[namespace] = &Resource{
			quota:          quota,
			metrics:        metrics,
			rejectedTokens: sync.Map{},
		}
		logger.Infof(ctx, "Creating namespacedResourcesMap: added namespace [%v] and resource [%v]", namespace, rm.namespacedResourcesMap[namespace])
	}
	rm.startMetricsGathering(ctx)
	return rm, nil
}

func NewRedisResourceManagerBuilder(_ context.Context, client RedisClient, scope promutils.Scope) (*RedisResourceManagerBuilder, error) {
	rn := &RedisResourceManagerBuilder{
		client:                      client,
		MetricsScope:                scope,
		namespacedResourcesQuotaMap: map[pluginCore.ResourceNamespace]int{},
	}

	return rn, nil
}

type RedisResourceManager struct {
	client                 RedisClient
	MetricsScope           promutils.Scope
	namespacedResourcesMap map[pluginCore.ResourceNamespace]*Resource
}

func GetTaskResourceManager(r pluginCore.ResourceManager, resourceNamespacePrefix pluginCore.ResourceNamespace, allocationTokenPrefix TokenPrefix) pluginCore.ResourceManager {
	return Proxy{
		ResourceManager:         r,
		ResourceNamespacePrefix: resourceNamespacePrefix,
		TokenPrefix:             allocationTokenPrefix,
	}
}

type RedisResourceManagerMetrics struct {
	Scope                     promutils.Scope
	RedisSizeCheckTime        promutils.StopWatch
	AllocatedTokensGauge      prometheus.Gauge
	ApproximateBackedUpLength prometheus.Gauge
}

func (rrmm RedisResourceManagerMetrics) GetScope() promutils.Scope {
	return rrmm.Scope
}

func (r *RedisResourceManager) getResource(namespace pluginCore.ResourceNamespace) (*Resource, error) {
	if resource, ok := r.namespacedResourcesMap[namespace]; ok {
		return resource, nil
	}
	return nil, errors.Errorf("Requested resource [%v] not found in namespacedResourceMap", namespace)
}

func (r *RedisResourceManager) pollRedis(ctx context.Context, namespace pluginCore.ResourceNamespace) {
	resource, err := r.getResource(namespace)
	if err != nil {
		return
	}
	stopWatch := resource.metrics.(*RedisResourceManagerMetrics).RedisSizeCheckTime.Start()
	defer stopWatch.Stop()
	size, err := r.client.SCard(string(namespace)).Result()
	if err != nil {
		logger.Errorf(ctx, "Error getting size of Redis set in metrics poller %v", err)
		return
	}

	metrics := resource.metrics.(*RedisResourceManagerMetrics)
	metrics.AllocatedTokensGauge.Set(float64(size))
	rejectedTokensCount := 0
	resource.rejectedTokens.Range(func(key, value interface{}) bool {
		rejectedTokensCount++
		return true
	})
	metrics.ApproximateBackedUpLength.Set(float64(rejectedTokensCount))
}

func (r *RedisResourceManager) startMetricsGathering(ctx context.Context) {
	go wait.Until(func() {
		for namespace := range r.namespacedResourcesMap {
			r.pollRedis(ctx, namespace)
		}
	}, 10*time.Second, ctx.Done())
}

func NewRedisResourceManagerMetrics(scope promutils.Scope) *RedisResourceManagerMetrics {
	return &RedisResourceManagerMetrics{
		Scope: scope,
		RedisSizeCheckTime: scope.MustNewStopWatch("size_check_time",
			"The time it takes to measure the size of the Redis Set where all utilized resource are stored", time.Millisecond),

		AllocatedTokensGauge: scope.MustNewGauge("size",
			"The number of allocation resourceRegistryTokens currently in the Redis set"),

		ApproximateBackedUpLength: scope.MustNewGauge("approx_backup",
			"Approximation for how long the current not-fulfilled-tokens queue is."),
	}
}

func (r *RedisResourceManager) GetID() string {
	return RedisSetKeyPrefix
}

func (r *RedisResourceManager) AllocateResource(ctx context.Context, namespace pluginCore.ResourceNamespace, allocationToken string) (
	pluginCore.AllocationStatus, error) {

	namespacedResource, err := r.getResource(namespace)
	if err != nil {
		logger.Errorf(ctx, "Error finding resource [%v] during allocation", namespace)
		return pluginCore.AllocationUndefined, err
	}
	// Check to see if the allocation token is already in the set
	found, err := r.client.SIsMember(string(namespace), allocationToken).Result()
	if err != nil {
		logger.Errorf(ctx, "Error getting size of Redis set %v", err)
		return pluginCore.AllocationUndefined, err
	}
	if found {
		logger.Infof(ctx, "Already allocated [%s:%s]", namespace, allocationToken)
		return pluginCore.AllocationStatusGranted, nil
	}

	size, err := r.client.SCard(string(namespace)).Result()
	if err != nil {
		logger.Errorf(ctx, "Error getting size of Redis set %v", err)
		return pluginCore.AllocationUndefined, err
	}

	if size > int64(namespacedResource.quota) {
		logger.Infof(ctx, "Too many allocations (total [%d]), rejecting [%s:%s]", size, namespace, allocationToken)
		namespacedResource.rejectedTokens.Store(allocationToken, struct{}{})
		return pluginCore.AllocationStatusExhausted, nil
	}

	countAdded, err := r.client.SAdd(string(namespace), allocationToken).Result()
	if err != nil {
		logger.Errorf(ctx, "Error adding token [%s:%s] %v", namespace, allocationToken, err)
		return pluginCore.AllocationUndefined, err
	}

	logger.Infof(ctx, "Added %d to the Redis Qubole set", countAdded)
	namespacedResource.rejectedTokens.Delete(allocationToken)

	return pluginCore.AllocationStatusGranted, err
}

func (r *RedisResourceManager) ReleaseResource(ctx context.Context, namespace pluginCore.ResourceNamespace, allocationToken string) error {
	countRemoved, err := r.client.SRem(string(namespace), allocationToken).Result()
	if err != nil {
		logger.Errorf(ctx, "Error removing token [%v:%s] %v", namespace, allocationToken, err)
		return err
	}

	namespacedResource, err := r.getResource(namespace)
	if err != nil {
		logger.Errorf(ctx, "Error finding resource [%v] during releasing", namespace)
		return err
	}
	namespacedResource.rejectedTokens.Delete(allocationToken)
	logger.Infof(ctx, "Removed %d token: %s", countRemoved, allocationToken)

	return nil
}
