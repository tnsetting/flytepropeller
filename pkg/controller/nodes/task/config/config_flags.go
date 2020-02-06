// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "task-plugins.enabled-plugins"), []string{}, "Plugins enabled currently")
	cmdFlags.Int32(fmt.Sprintf("%v%v", prefix, "max-plugin-phase-versions"), defaultConfig.MaxPluginPhaseVersions, "Maximum number of plugin phase versions allowed for one phase.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "barrier.enabled"), defaultConfig.BarrierConfig.Enabled, "Enable Barrier transitions using inmemory context")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "barrier.cache-size"), defaultConfig.BarrierConfig.CacheSize, "Max number of barrier to preserve in memory")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "barrier.cache-ttl"), defaultConfig.BarrierConfig.CacheTTL.String(), " Max duration that a barrier would be respected if the process is not restarted. This should account for time required to store the record into persistent storage (across multiple rounds.")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "backoff.base-second"), defaultConfig.BackOffConfig.BaseSecond, "The number of seconds representing the base duration of the exponential backoff")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "backoff.max-duration"), defaultConfig.BackOffConfig.MaxDuration.String(), "The cap of the backoff duration")
	return cmdFlags
}
