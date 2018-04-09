package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Map
type Map map[string]interface{}

// Exist judge if the key exist in map
func (m Map) Exist(key string) bool {
	if _, ok := m[key]; !ok {
		return false
	}
	return true
}

// String return the value type string by key
func (m Map) String(key string) string {
	var val string
	var ok bool
	if val, ok = m[key].(string); ok {
		return val
	}

	res, _ := TransType(m[key], reflect.String)
	if val, ok = res.(string); !ok {
		return ""
	}
	return val
}

// Int return the value type int by key
func (m Map) Int(key string) int {
	var val int
	var ok bool
	if val, ok = m[key].(int); ok {
		return val
	}
	res, _ := TransType(m[key], reflect.Int)
	if val, ok = res.(int); !ok {
		return 0
	}
	return val
}

// Int return the value type int32 by key
func (m Map) Int32(key string) int32 {
	var val int32
	var ok bool
	if val, ok = m[key].(int32); ok {
		return val
	}
	res, _ := TransType(m[key], reflect.Int32)
	if val, ok = res.(int32); !ok {
		return 0
	}
	return val
}

// Int return the value type int64 by key
func (m Map) Int64(key string) int64 {
	var val int64
	var ok bool
	if val, ok = m[key].(int64); ok {
		return val
	}
	res, _ := TransType(m[key], reflect.Int64)
	if val, ok = res.(int64); !ok {
		return 0
	}
	return val
}

// Int return the value type float32 by key
func (m Map) Float32(key string) float32 {
	var val float32
	var ok bool
	if val, ok = m[key].(float32); ok {
		return val
	}
	res, _ := TransType(m[key], reflect.Float32)
	if val, ok = res.(float32); !ok {
		return 0
	}
	return val
}

// Int return the value type float32 by key
func (m Map) Float64(key string) float64 {
	var val float64
	var ok bool
	if val, ok = m[key].(float64); ok {
		return val
	}
	res, _ := TransType(m[key], reflect.Float64)
	if val, ok = res.(float64); !ok {
		return 0
	}
	return val
}

// Int return the value type Map by key
func (m Map) Map(key string) Map {
	var val Map
	var ok bool
	if val, ok = m[key].(Map); ok {
		return val
	} else if val, ok = m[key].(map[string]interface{}); ok {
		return Map(val)
	}
	return nil
}

// S2Json trans data to json, e.g struct, map and so on.
func S2Json(data interface{}) string {
	bys, _ := json.Marshal(data)
	return string(bys)
}

// Json2S trans json to object
func Json2S(src string, dest interface{}) error {
	return json.Unmarshal([]byte(src), dest)
}

func TransType(val interface{}, descType reflect.Kind) (interface{}, error) {
	if val == nil {
		return nil, errors.New("val is nil")
	}
	typ := reflect.TypeOf(val).Kind()
	switch typ {
	case reflect.Int:
		res, _ := val.(int)
		switch descType {
		case reflect.Int:
			return int(res), nil
		case reflect.Int32:
			return int32(res), nil
		case reflect.Int64:
			return int64(res), nil
		case reflect.Float32:
			return float32(res), nil
		case reflect.Float64:
			return float64(res), nil
		case reflect.String:
			return fmt.Sprintf("%v", res), nil
		}
	case reflect.Int32:
		res, _ := val.(int32)
		switch descType {
		case reflect.Int:
			return int(res), nil
		case reflect.Int32:
			return int32(res), nil
		case reflect.Int64:
			return int64(res), nil
		case reflect.Float32:
			return float32(res), nil
		case reflect.Float64:
			return float64(res), nil
		case reflect.String:
			return fmt.Sprintf("%v", res), nil
		}
	case reflect.Int64:
		res, _ := val.(int64)
		switch descType {
		case reflect.Int:
			return int(res), nil
		case reflect.Int32:
			return int32(res), nil
		case reflect.Int64:
			return int64(res), nil
		case reflect.Float32:
			return float32(res), nil
		case reflect.Float64:
			return float64(res), nil
		case reflect.String:
			return fmt.Sprintf("%v", res), nil
		}
	case reflect.Float32:
		res, _ := val.(float32)
		switch descType {
		case reflect.Int:
			return int(res), nil
		case reflect.Int32:
			return int32(res), nil
		case reflect.Int64:
			return int64(res), nil
		case reflect.Float32:
			return float32(res), nil
		case reflect.Float64:
			return float64(res), nil
		case reflect.String:
			return fmt.Sprintf("%v", res), nil
		}
	case reflect.Float64:
		res, _ := val.(float64)
		switch descType {
		case reflect.Int:
			return int(res), nil
		case reflect.Int32:
			return int32(res), nil
		case reflect.Int64:
			return int64(res), nil
		case reflect.Float32:
			return float32(res), nil
		case reflect.Float64:
			return float64(res), nil
		case reflect.String:
			return fmt.Sprintf("%v", res), nil
		}
	case reflect.String:
		res, _ := val.(string)
		switch descType {
		case reflect.Int:
			s_res, _ := strconv.Atoi(res)
			return s_res, nil
		case reflect.Int32:
			s_res, _ := strconv.Atoi(res)
			return int32(s_res), nil
		case reflect.Int64:
			s_res, _ := strconv.Atoi(res)
			return int64(s_res), nil
		case reflect.Float32:
			s_res, _ := strconv.ParseFloat(res, 32)
			return float32(s_res), nil
		case reflect.Float64:
			s_res, _ := strconv.ParseFloat(res, 64)
			return float64(s_res), nil
		case reflect.String:
			return fmt.Sprintf("%v", res), nil
		}
	}
	return nil, errors.New(fmt.Sprintf("invalid value type(%v)", typ))
}
