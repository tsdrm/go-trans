package util

import (
	"fmt"
	"math"
	"testing"
	"w.gdy.io/dyf/transcode/util"
)

func TestMap(t *testing.T) {
	m := Map{}

	// 检查语法错误
	m.String("halo")
	m.Int("halo")
	m.Int32("halo")
	m.Int64("halo")
	m.Float32("halo")
	m.Float64("halo")
	m.Map("halo")

	// 检查值
	m = Map{
		"string":  "string",
		"int":     1,
		"int32":   int32(2),
		"int64":   int64(3),
		"float32": float32(1.1),
		"float64": float64(2.2),
		"Map":     Map{"key": "val"},
	}
	if "string" != m.String("string") {
		t.Error(m.String("string"))
		return
	}
	if 1 != m.Int("int") {
		t.Error(m.Int("int"))
		return
	}
	if 2 != m.Int32("int32") {
		t.Error(m.Int32("int32"))
		return
	}
	if 3 != m.Int64("int64") {
		t.Error(m.Int64("int64"))
		return
	}
	if 1.1 != m.Float32("float32") {
		t.Error(m.Float32("float32"))
		return
	}
	if 2.2 != m.Float64("float64") {
		t.Error("float64")
		return
	}
	if nil == m.Map("Map") {
		t.Error(m.Map("Map"))
		return
	}
	mm := m.Map("Map")
	if "val" != mm.String("key") {
		t.Error(mm.String("key"))
		return
	}
	if true != m.Exist("string") {
		t.Error(m.Exist("string"))
	}
	if false != m.Exist("halo") {
		t.Error(m.Exist("halo"))
	}

	check(t, m)

	m = Map{
		"string":  "string",
		"int":     "1",
		"int32":   "2",
		"int64":   "3",
		"float32": "1.1",
		"float64": "2.2",
		"Map":     Map{"key": "val"},
	}

	// int
	if 1 != m.Int("int") {
		t.Error(m.Int("int"))
		return
	}
	if 2 != m.Int("int32") {
		t.Error(m.Int("int32"))
		return
	}
	if 3 != m.Int("int64") {
		t.Error(m.Int("int64"))
		return
	}
	if 0 != m.Int("float32") {
		t.Error(m.Int("float32"))
		return
	}
	if 0 != m.Int("float64") {
		t.Error(m.Int("float64"))
		return
	}
	if 0 != m.Int("string") {
		t.Error(m.Int("string"))
		return
	}

	// int32
	if 1 != m.Int32("int") {
		t.Error(m.Int32("int"))
		return
	}
	if 2 != m.Int32("int32") {
		t.Error(m.Int32("int32"))
		return
	}
	if 3 != m.Int32("int64") {
		t.Error(m.Int32("int64"))
		return
	}
	if 0 != m.Int32("float32") {
		t.Error(m.Int32("float32"))
		return
	}
	if 0 != m.Int32("float64") {
		t.Error(m.Int32("float64"))
		return
	}
	if 0 != m.Int32("string") {
		t.Error(m.Int32("string"))
		return
	}

	// int64
	if 1 != m.Int64("int") {
		t.Error(m.Int64("int"))
		return
	}
	if 2 != m.Int64("int32") {
		t.Error(m.Int64("int32"))
		return
	}
	if 3 != m.Int64("int64") {
		t.Error(m.Int64("int64"))
		return
	}
	if 0 != m.Int64("float32") {
		t.Error(m.Int64("float32"))
		return
	}
	if 0 != m.Int64("float64") {
		t.Error(m.Int64("float64"))
		return
	}
	if 0 != m.Int64("string") {
		t.Error(m.Int64("string"))
		return
	}

	// float32
	if 1 != m.Float32("int") {
		t.Error(m.Float32("int"))
		return
	}
	if 2 != m.Float32("int32") {
		t.Error(m.Float32("int32"))
		return
	}
	if 3 != m.Float32("int64") {
		t.Error(m.Float32("int64"))
		return
	}
	if 1.1 != m.Float32("float32") {
		t.Error(m.Float32("float32"))
		return
	}
	if 2.2 != m.Float32("float64") {
		t.Error(m.Float32("float64"))
		return
	}
	if 0 != m.Float32("string") {
		t.Error(m.Float32("string"))
		return
	}

	// float64
	if 1 != m.Float64("int") {
		t.Error(m.Float64("int"))
		return
	}
	if 2 != m.Float64("int32") {
		t.Error(m.Float64("int32"))
		return
	}
	if 3 != m.Float64("int64") {
		t.Error(m.Float64("int64"))
		return
	}
	//if math.Abs(float64(1.1 - m.Float64("float32"))) > 0.0001 {
	//	t.Error(m.Float64("float32"))
	//	return
	//}
	if 1.1 != m.Float64("float32") {
		t.Error(m.Float32("float32"))
		return
	}
	if 2.2 != m.Float64("float64") {
		t.Error(m.Float64("float64"))
		return
	}
	if 0 != m.Float64("string") {
		t.Error(m.Float64("string"))
		return
	}

	// string
	if "1" != m.String("int") {
		t.Error(m.String("int"))
		return
	}
	if "2" != m.String("int32") {
		t.Error(m.String("int32"))
		return
	}
	if "3" != m.String("int64") {
		t.Error(m.String("int64"))
		return
	}
	if "1.1" != m.String("float32") {
		t.Error(m.String("float32"))
		return
	}
	if "2.2" != m.String("float64") {
		t.Error(m.String("float64"))
		return
	}
	if "string" != m.String("string") {
		t.Error(m.String("string"))
		return
	}

	// check map
	m = Map{
		"a": Map{
			"1": 2,
		},
		"b": Map{
			"2": "2",
		},
	}
	a := m.Map("a")
	if a == nil {
		t.Error(nil)
		return
	}
	if 2 != a.Int("1") {
		t.Error(a.Int("1"))
		return
	}
	b := m.Map("b")
	if b == nil {
		t.Error(nil)
		return
	}
	if "2" != b.String("2") {
		t.Error(b.String("2"))
		return
	}
	fmt.Println("TestMap test success")
}

func check(t *testing.T, m Map) {
	// int
	if 2 != m.Int("int32") {
		t.Error(m.Int("int32"))
		return
	}
	if 3 != m.Int("int64") {
		t.Error(m.Int("int64"))
		return
	}
	if 1 != m.Int("float32") {
		t.Error(m.Int("float32"))
		return
	}
	if 2 != m.Int("float64") {
		t.Error(m.Int("float64"))
		return
	}
	if 0 != m.Int("string") {
		t.Error(m.Int("string"))
		return
	}

	// int32
	if 1 != m.Int32("int") {
		t.Error(m.Int32("int"))
		return
	}
	if 3 != m.Int32("int64") {
		t.Error(m.Int32("int64"))
		return
	}
	if 1 != m.Int32("float32") {
		t.Error(m.Int32("float32"))
		return
	}
	if 2 != m.Int32("float64") {
		t.Error(m.Int32("float64"))
		return
	}
	if 0 != m.Int32("string") {
		t.Error(m.Int32("string"))
		return
	}

	// int64
	if 1 != m.Int64("int") {
		t.Error(m.Int64("int"))
		return
	}
	if 2 != m.Int64("int32") {
		t.Error(m.Int64("int32"))
		return
	}
	if 1 != m.Int64("float32") {
		t.Error(m.Int64("float32"))
		return
	}
	if 2 != m.Int64("float64") {
		t.Error(m.Int64("float64"))
		return
	}
	if 0 != m.Int64("string") {
		t.Error(m.Int64("string"))
		return
	}

	// float32
	if 1 != m.Float32("int") {
		t.Error(m.Float32("int"))
		return
	}
	if 2 != m.Float32("int32") {
		t.Error(m.Float32("int32"))
		return
	}
	if 3 != m.Float32("int64") {
		t.Error(m.Float32("int64"))
		return
	}
	if 2.2 != m.Float32("float64") {
		t.Error(m.Float32("float64"))
		return
	}
	if 0 != m.Float32("string") {
		t.Error(m.Float32("string"))
		return
	}

	// float64
	if 1 != m.Float64("int") {
		t.Error(m.Float64("int"))
		return
	}
	if 2 != m.Float64("int32") {
		t.Error(m.Float64("int32"))
		return
	}
	if 3 != m.Float64("int64") {
		t.Error(m.Float64("int64"))
		return
	}
	if math.Abs(float64(1.1-m.Float64("float32"))) > 0.0001 {
		t.Error(m.Float64("float32"))
		return
	}
	if 0 != m.Float64("string") {
		t.Error(m.Float64("string"))
		return
	}

	// string
	if "1" != m.String("int") {
		t.Error(m.String("int"))
		return
	}
	if "2" != m.String("int32") {
		t.Error(m.String("int32"))
		return
	}
	if "3" != m.String("int64") {
		t.Error(m.String("int64"))
		return
	}
	if "1.1" != m.String("float32") {
		t.Error(m.String("float32"))
		return
	}
	if "2.2" != m.String("float64") {
		t.Error(m.String("float64"))
		return
	}
}

func TestAry(t *testing.T) {
	var data = Map{
		"aryMap": []Map{
			{"a": "b"},
		},
	}
	var vals = data.AryMap("aryMap")
	if len(vals) != 1 {
		t.Error(len(vals))
		return
	}
	if vals[0].String("a") != "b" {
		t.Error(vals[0])
		return
	}

	//
	var str = util.S2Json(data)
	var err = Json2S(str, &data)
	if err != nil {
		t.Error(err)
		return
	}

	if len(vals) != 1 {
		t.Error(len(vals))
		return
	}
	if vals[0].String("a") != "b" {
		t.Error(vals[0])
		return
	}
}
