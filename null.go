package jsontype

import (
	"encoding/json"
)

type NullType interface {
	IsNull() bool
}

type NullString struct {
	Null  bool
	Value string
}

func (ns NullString) IsNull() bool {
	return ns.Null
}

func (ns NullString) String() string {
	if ns.Null {
		return ""
	}
	return ns.Value
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.IsNull() {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(ns.Value))
	b = append(b, '"')
	b = append(b, []byte(ns.Value)...)
	b = append(b, '"')
	return b, nil
}

func (ns *NullString) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		ns.Null = true
		ns.Value = ""
		return nil
	} else {
		var str string
		err := json.Unmarshal(b, &str)
		if err != nil {
			return err
		}
		ns.Null = false
		ns.Value = str
		return nil
	}
}

type NullInt struct {
	Null  bool
	Value int
}

func (ns NullInt) IsNull() bool {
	return ns.Null
}

func (ns NullInt) Int() int {
	if ns.Null {
		return 0
	}
	return ns.Value
}

func (ns NullInt) MarshalJSON() ([]byte, error) {
	if ns.IsNull() {
		return []byte("null"), nil
	}

	return json.Marshal(ns.Value)
}

func (ns *NullInt) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		ns.Null = true
		ns.Value = 0
		return nil
	} else {
		var i int
		err := json.Unmarshal(b, &i)
		if err != nil {
			return err
		}
		ns.Null = false
		ns.Value = i
		return nil
	}
}

type NullBool struct {
	Null  bool
	Value bool
}

func (ns NullBool) IsNull() bool {
	return ns.Null
}

func (ns NullBool) Bool() bool {
	if ns.Null {
		return false
	}
	return ns.Value
}

func (ns NullBool) MarshalJSON() ([]byte, error) {
	if ns.IsNull() {
		return []byte("null"), nil
	}

	return json.Marshal(ns.Value)
}

func (ns *NullBool) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		ns.Null = true
		ns.Value = false
		return nil
	} else {
		var boolean bool
		err := json.Unmarshal(b, &boolean)
		if err != nil {
			return err
		}
		ns.Null = false
		ns.Value = boolean
		return nil
	}
}
