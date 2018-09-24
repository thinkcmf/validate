package validate

import (
	"fmt"
	"reflect"
	"strings"
)

// CallByValue call func by reflect.Value
func CallByValue(fv reflect.Value, args ...interface{}) []reflect.Value {
	if fv.Kind() != reflect.Func {
		panic("parameter must be an func type")
	}

	argNum := len(args)
	if argNum < fv.Type().NumIn() {
		fmt.Println("the number of input params not match!")
	}

	in := make([]reflect.Value, argNum)
	for k, v := range args {
		in[k] = reflect.ValueOf(v)
	}

	// CallSlice()与Call() 不一样的是，参数的最后一个会被展开
	// f.CallSlice()
	return fv.Call(in)
}

// Call call func by reflection
func Call(fn interface{}, args ...interface{}) []reflect.Value {
	return CallByValue(reflect.ValueOf(fn), args...)
}

// upperFirst upper first char
func upperFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	f := s[0]

	if f >= 'a' && f <= 'z' {
		return strings.ToUpper(string(f)) + string(s[1:])
	}

	return s
}

// Length calc
func Length(val interface{}) int {
	if val == nil {
		return -1
	}

	if rv, ok := val.(reflect.Value); !ok {
		return ValueLen(rv)
	}

	return ValueLen(reflect.ValueOf(val))
}

// IntVal of the val
func IntVal(val interface{}) (intVal int64, ok bool) {
	switch tv := val.(type) {
	case int:
		ok = true
		intVal = int64(tv)
	case int64:
		ok = true
		intVal = tv
	case reflect.Value:
		intVal, ok = ValueInt64(tv)
	default:
		intVal, ok = ValueInt64(reflect.ValueOf(val))
	}

	return
}

func int64compare(intVal, dstVal int64, op string) bool {
	switch op {
	case "eq":
		return intVal == dstVal
	case "ne":
		return intVal != dstVal
	case "lt":
		return intVal < dstVal
	case "lte":
		return intVal <= dstVal
	case "gt":
		return intVal > dstVal
	case "gte":
		return intVal >= dstVal
	}

	return false
}
