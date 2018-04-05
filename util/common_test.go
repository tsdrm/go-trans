package util

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
)

func TestUuid(t *testing.T) {
	var times int = 10
	for i := 0; i < times; i++ {
		testUuid(t, times)
	}

	log.Println("TestUuid test success")
}

// TrimAryStringRepeat trim the s in every string in strs,
// repeat is decide result string array whether allowed repeat,
// false is don't allowed
func trimAryStringRepeat(strs []string, s string, repeat bool) []string {
	var strMap = map[string]bool{}
	var results = []string{}
	for _, str := range strs {
		val := strings.Trim(str, s)
		if len(val) < 1 {
			continue
		}
		if strMap[val] && !repeat {
			continue
		}
		strMap[val] = true
		results = append(results, val)
	}
	return results
}

func testUuid(t *testing.T, times int) {
	fmt.Println("times --- ", times, " START")
	uuid := UUID()
	if "" == uuid {
		t.Error(uuid)
		return
	}
	t.Log(uuid)

	// 连续生成uuid 10000个,查看是否有重复的id生成
	var length = 10000
	strs := make([]string, 0, length+10)
	for i := 0; i < length; i++ {
		strs = append(strs, UUID())
	}
	strs = trimAryStringRepeat(strs, "", false)
	if len(strs) != length {
		t.Error(len(strs))
	}

	// 并发生成uuid 10000个,查看是否有重复的uuid生成
	strs = []string{}
	strs = make([]string, 0, length+100)
	for i := 0; i < length; i++ {
		go func() {
			strs = append(strs, UUID())
		}()
	}
	time.Sleep(time.Second * 1)
	originLength := len(strs)
	strs = trimAryStringRepeat(strs, "", false)
	if len(strs) != originLength {
		t.Error(len(strs))
	}
	fmt.Println("times --- ", times, " END")
}

func TestError(t *testing.T) {
	var err error
	err = NewError("%v", "halo")

	if err == nil || err.Error() != "halo" {
		t.Error(err)
		return
	}

	err = NewError("aaa%vbbb%v", 123, 456)
	if err == nil || err.Error() != "aaa123bbb456" {
		t.Error(err)
		return
	}

	log.Println("TestError test success")
}

func TestNow(t *testing.T) {
	var now = Now13()
	fmt.Println(now)
	if now/1e12 > 10 || now/1e12 < 1 {
		t.Error(now)
		return
	}

	now = Now10()
	fmt.Println(now)
	if now/1e9 > 10 || now/1e9 < 1 {
		t.Error(now)
		return
	}
}

func TestS2Json(t *testing.T) {
	var output string
	var mp = map[string]string{
		"a": "b",
	}
	output = S2Json(mp)
	if output != `{"a":"b"}` {
		t.Error(output)
		return
	}

	var mp_check map[string]string
	var err = Json2S(output, &mp_check)
	if err != nil {
		t.Error(err)
		return
	}
	if mp_check == nil || mp_check["a"] != "b" || len(mp_check) != 1 {
		t.Error(mp_check)
		return
	}

	err = Json2S("halo", &mp_check)
	if err == nil {
		t.Error(err)
		return
	}
}
