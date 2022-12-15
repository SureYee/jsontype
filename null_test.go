package jsontype_test

import (
	"encoding/json"
	"testing"

	"github.com/sureyee/jsontype"
)

func TestNullObject(t *testing.T) {
	var jsonStr = []byte(`{"a":{"a":"b"}, "n":null}`)
	var obj jsontype.Object
	if err := json.Unmarshal(jsonStr, &obj); err != nil {
		t.Error(err)
		return
	}
	newObj, err := obj.GetObject("a")
	if err != nil {
		t.Error(err)
		return
	}
	if newObj.IsNull() {
		t.Errorf("a的值不是null，但是判断为null")
		return
	}
	nullObj, err := obj.GetObject("n")
	if err != nil {
		t.Error(err)
		return
	}
	if !nullObj.IsNull() {
		t.Errorf("n的值是null，但是判断不为null")
		return
	}
}

func TestNullString(t *testing.T) {
	var jsonStr = []byte(`{"a": null, "b": "b"}`)
	var m map[string]jsontype.NullString
	err := json.Unmarshal(jsonStr, &m)
	if err != nil {
		t.Error(err)
	}

	if !m["a"].IsNull() {
		t.Errorf("a is null but check result is not null")
	}

	if m["b"].IsNull() {
		t.Errorf("b is not null but check result is null")
	}

	if m["b"].String() != "b" {
		t.Errorf("b value is %s but got %s", "b", m["b"].String())
	}
}

func TestNullInt(t *testing.T) {
	var jsonStr = []byte(`{"a": null, "b": 10}`)
	var m map[string]jsontype.NullInt
	err := json.Unmarshal(jsonStr, &m)
	if err != nil {
		t.Error(err)
	}

	if !m["a"].IsNull() {
		t.Errorf("a is null but check result is not null")
	}

	if m["b"].IsNull() {
		t.Errorf("b is not null but check result is null")
	}

	if m["b"].Int() != 10 {
		t.Errorf("b value is %d but got %d", 10, m["b"].Int())
	}
}

func TestNullBool(t *testing.T) {
	var jsonStr = []byte(`{"a": null, "b": true}`)
	var m map[string]jsontype.NullBool
	err := json.Unmarshal(jsonStr, &m)
	if err != nil {
		t.Error(err)
	}

	if !m["a"].IsNull() {
		t.Errorf("a is null but check result is not null")
	}

	if m["b"].IsNull() {
		t.Errorf("b is not null but check result is null")
	}

	if !m["b"].Bool() {
		t.Errorf("b value is %v but got %v", true, m["b"].Bool())
	}
}

func TestNullTypeMarshal(t *testing.T) {
	var a = make(map[string]any)
	a["a"] = jsontype.NullBool{Null: true}
	a["b"] = jsontype.NullBool{Null: false, Value: true}
	a["c"] = jsontype.NullInt{Null: true, Value: 10}
	a["d"] = jsontype.NullInt{Null: false, Value: 20}
	a["e"] = jsontype.NullString{Null: true, Value: "dsfaf"}
	a["f"] = jsontype.NullString{Null: false, Value: "dfsefef"}
	b, err := json.Marshal(a)
	t.Log(string(b), err)
}
