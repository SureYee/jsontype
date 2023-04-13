package jsontype_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/sureyee/jsontype"
)

var jsonObj = `{"obj":{"a":"b", "b": {"b1": "aaa"}}, "int": 10, "num": "10000", "str":"nnn", "arr": [1,2,3], "b": true, "n": null}`
var obj jsontype.Object

func TestObject(t *testing.T) {
	_, err := obj.GetObject("obj")
	if err != nil {
		t.Error(err)
		return
	}
	_, err = obj.GetString("str")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestObjectMarshal(t *testing.T) {
	var newObj = jsontype.NewObject()
	newObj.Set("int", 10)
	newObj.Set("obj", map[string]string{"a": "b"})
	newObj.Set("str", "nnn")
	newObj.Set("arr", []int{1, 2, 3})

	b, err := json.Marshal(newObj)
	if err != nil {
		t.Error(err)
		return

	}
	t.Log(string(b))
}

func TestGetInt(t *testing.T) {
	i, err := obj.GetInt("int")
	if err != nil {
		t.Error(err)
		return
	}

	if i != 10 {
		t.Errorf("预期值为[%d],但是获取到[%d]", 10, i)
		return
	}

	_, err = obj.GetInt("str")
	if !errors.Is(err, jsontype.ErrTypeNotMatch) {
		t.Errorf("预期错误为[%s],但是错误为[%s]", jsontype.ErrTypeNotMatch, err)
	}
}

func TestGetBool(t *testing.T) {
	value, err := obj.GetBool("b")
	if err != nil {
		t.Error(err)
		return
	}

	if !value {
		t.Errorf("预期值为[%v],但是获取到[%v]", true, value)
		return
	}
}

func TestGetInts(t *testing.T) {
	value, err := obj.GetInts("arr")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(value)
}

func TestGetNullString(t *testing.T) {
	value, err := obj.GetNullString("str")
	if err != nil {
		t.Error(err)
		return
	}

	if value.IsNull() {
		t.Error("value is not null but got null")
	}

	if value.String() != "nnn" {
		t.Errorf("value is %s but got %s", "nnn", value.String())
	}
}

func TestGetNullInt(t *testing.T) {
	value, err := obj.GetNullInt("int")
	if err != nil {
		t.Error(err)
		return
	}

	if value.IsNull() {
		t.Error("value is not null but got null")
		return
	}

	if value.Int() != 10 {
		t.Errorf("value is %d but got %d", 10, value.Int())
	}
}

func TestGetNullBool(t *testing.T) {
	value, err := obj.GetNullBool("b")
	if err != nil {
		t.Error(err)
		return
	}

	if value.IsNull() {
		t.Error("value is not null but got null")
		return
	}

	if !value.Bool() {
		t.Errorf("value is %v but got %v", 10, value.Bool())
	}
}

func TestGetNumber(t *testing.T) {
	// test int number
	num, err := obj.GetNumber("int")
	if err != nil {
		t.Error(err)
		return
	}
	if i, err := num.Int64(); err != nil {
		t.Error(err)
		return
	} else {
		if i != 10 {
			t.Errorf("期望值为[%v],但是获取到值为[%v]", 10, i)
		}
	}

}

func TestSet(t *testing.T) {
	o, _ := obj.GetObject("obj")
	o.Set("c", "c")
	b, _ := json.Marshal(obj)
	fmt.Printf("%s\n", b)
}

type OtherObject struct {
	jsontype.Object
}

func TestStruct(t *testing.T) {
	var obj OtherObject
	err := json.Unmarshal([]byte(jsonObj), &obj)
	if err != nil {
		t.Error(err)
	}
	t.Log(obj)
}

func TestStructSet(t *testing.T) {
	var obj OtherObject
	err := json.Unmarshal([]byte(jsonObj), &obj)
	o, _ := obj.GetObject("obj")
	bo, _ := o.GetObject("b")
	bo.Set("b2", "newB2")
	b, _ := json.Marshal(obj)
	fmt.Printf("%s\n", b)
	if err != nil {
		t.Error(err)
	}
	t.Log(obj)
}

func init() {
	err := json.Unmarshal([]byte(jsonObj), &obj)
	if err != nil {
		panic(err)
	}
}
