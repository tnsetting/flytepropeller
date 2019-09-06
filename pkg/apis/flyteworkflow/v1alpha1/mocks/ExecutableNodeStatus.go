// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
import mock "github.com/stretchr/testify/mock"
import storage "github.com/lyft/flytestdlib/storage"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"

// ExecutableNodeStatus is an autogenerated mock type for the ExecutableNodeStatus type
type ExecutableNodeStatus struct {
	mock.Mock
}

// ClearDynamicNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ClearDynamicNodeStatus() {
	_m.Called()
}

// ClearSubWorkflowStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ClearSubWorkflowStatus() {
	_m.Called()
}

// ClearTaskStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ClearTaskStatus() {
	_m.Called()
}

// ClearWorkflowStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ClearWorkflowStatus() {
	_m.Called()
}

// GetAttempts provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetAttempts() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// GetDataDir provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetDataDir() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

// GetLastUpdatedAt provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetLastUpdatedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

// GetMessage provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetMessage() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNodeExecutionStatus provides a mock function with given fields: id
func (_m *ExecutableNodeStatus) GetNodeExecutionStatus(id string) v1alpha1.ExecutableNodeStatus {
	ret := _m.Called(id)

	var r0 v1alpha1.ExecutableNodeStatus
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableNodeStatus); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNodeStatus)
		}
	}

	return r0
}

// GetOrCreateBranchStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetOrCreateBranchStatus() v1alpha1.MutableBranchNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableBranchNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableBranchNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableBranchNodeStatus)
		}
	}

	return r0
}

// GetOrCreateDynamicNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetOrCreateDynamicNodeStatus() v1alpha1.MutableDynamicNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableDynamicNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableDynamicNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableDynamicNodeStatus)
		}
	}

	return r0
}

// GetOrCreateSubWorkflowStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetOrCreateSubWorkflowStatus() v1alpha1.MutableSubWorkflowNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableSubWorkflowNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableSubWorkflowNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableSubWorkflowNodeStatus)
		}
	}

	return r0
}

// GetOrCreateTaskStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetOrCreateTaskStatus() v1alpha1.MutableTaskNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableTaskNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableTaskNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableTaskNodeStatus)
		}
	}

	return r0
}

// GetOrCreateWorkflowStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetOrCreateWorkflowStatus() v1alpha1.MutableWorkflowNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableWorkflowNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableWorkflowNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableWorkflowNodeStatus)
		}
	}

	return r0
}

// GetParentNodeID provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetParentNodeID() *string {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// GetParentTaskID provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetParentTaskID() *core.TaskExecutionIdentifier {
	ret := _m.Called()

	var r0 *core.TaskExecutionIdentifier
	if rf, ok := ret.Get(0).(func() *core.TaskExecutionIdentifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TaskExecutionIdentifier)
		}
	}

	return r0
}

// GetPhase provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetPhase() v1alpha1.NodePhase {
	ret := _m.Called()

	var r0 v1alpha1.NodePhase
	if rf, ok := ret.Get(0).(func() v1alpha1.NodePhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.NodePhase)
	}

	return r0
}

// GetQueuedAt provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetQueuedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

// GetStartedAt provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetStartedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

// GetStoppedAt provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetStoppedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

// GetSubWorkflowNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetSubWorkflowNodeStatus() v1alpha1.ExecutableSubWorkflowNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableSubWorkflowNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableSubWorkflowNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableSubWorkflowNodeStatus)
		}
	}

	return r0
}

// GetTaskNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetTaskNodeStatus() v1alpha1.ExecutableTaskNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableTaskNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableTaskNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableTaskNodeStatus)
		}
	}

	return r0
}

// GetWorkflowNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetWorkflowNodeStatus() v1alpha1.ExecutableWorkflowNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableWorkflowNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableWorkflowNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableWorkflowNodeStatus)
		}
	}

	return r0
}

// IncrementAttempts provides a mock function with given fields:
func (_m *ExecutableNodeStatus) IncrementAttempts() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// IsCached provides a mock function with given fields:
func (_m *ExecutableNodeStatus) IsCached() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsDirty provides a mock function with given fields:
func (_m *ExecutableNodeStatus) IsDirty() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ResetDirty provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ResetDirty() {
	_m.Called()
}

// SetCached provides a mock function with given fields:
func (_m *ExecutableNodeStatus) SetCached() {
	_m.Called()
}

// SetDataDir provides a mock function with given fields: _a0
func (_m *ExecutableNodeStatus) SetDataDir(_a0 storage.DataReference) {
	_m.Called(_a0)
}

// SetParentNodeID provides a mock function with given fields: n
func (_m *ExecutableNodeStatus) SetParentNodeID(n *string) {
	_m.Called(n)
}

// SetParentTaskID provides a mock function with given fields: t
func (_m *ExecutableNodeStatus) SetParentTaskID(t *core.TaskExecutionIdentifier) {
	_m.Called(t)
}

// UpdatePhase provides a mock function with given fields: phase, occurredAt, reason
func (_m *ExecutableNodeStatus) UpdatePhase(phase v1alpha1.NodePhase, occurredAt v1.Time, reason string) {
	_m.Called(phase, occurredAt, reason)
}

// VisitNodeStatuses provides a mock function with given fields: visitor
func (_m *ExecutableNodeStatus) VisitNodeStatuses(visitor func(string, v1alpha1.ExecutableNodeStatus)) {
	_m.Called(visitor)
}