package job

import (
	"reflect"
	"runtime"
)

// Status status
func Status() map[string]string {
	val := make(map[string]string)
	for k, v := range handlers {
		val[k] = runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
	}
	return val
}
