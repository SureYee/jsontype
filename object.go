package jsontype

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"sync"
)

var ErrKeyNotExists = errors.New("key not exists")
var ErrTypeNotMatch = errors.New("type not match")

type Object struct {
	locker *sync.Mutex
	raw    map[string]json.RawMessage // 存放json的字节
	attrs  map[string]any             // 存放获取的属性，如果没有获取过，则是空的，序列化的时候用属性替换掉raw中的数据
}

// IsNull
// 判断object是否为null
func (obj *Object) IsNull() bool {
	return obj.raw == nil && len(obj.attrs) == 0
}

// GetString
// 获取key为string的值，如果不存在返回 ErrKeyNotExists
// 如果类型不正确，返回 ErrTypeNotMatch
func (obj *Object) GetString(key string) (string, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		if str, ok := attr.(string); ok {
			return str, nil
		}
		return "", ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return "", ErrKeyNotExists
	}
	var str string
	err := json.Unmarshal(raw, &str)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return "", ErrTypeNotMatch
		}
		return "", err
	}
	return str, nil
}

func (obj *Object) GetNullString(key string) (NullString, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		if str, ok := attr.(string); ok {
			return NullString{Null: false, Value: str}, nil
		}
		if nullStr, ok := attr.(NullString); ok {
			return nullStr, nil
		}

		if nullStr, ok := attr.(*NullString); ok {
			return *nullStr, nil
		}
		return NullString{}, ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return NullString{}, ErrKeyNotExists
	}
	var str NullString
	err := json.Unmarshal(raw, &str)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return str, ErrTypeNotMatch
		}
		return str, err
	}
	return str, nil
}

func (obj *Object) GetInt(key string) (int, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		if i, ok := attr.(int); ok {
			return i, nil
		}
		return 0, ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return 0, ErrKeyNotExists
	}
	var i int
	err := json.Unmarshal(raw, &i)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return 0, ErrTypeNotMatch
		}
		return 0, err
	}
	return i, nil
}

func (obj *Object) GetNumber(key string) (json.Number, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		v := reflect.ValueOf(attr)
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
			return json.Number(strconv.FormatInt(v.Int(), 10)), nil
		case reflect.String:
			return json.Number(v.String()), nil
		}
		return "", ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return "", ErrKeyNotExists
	}
	var value json.Number
	err := json.Unmarshal(raw, &value)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return "", ErrTypeNotMatch
		}
		return "", err
	}
	return value, nil
}

func (obj *Object) GetNumbers(key string) ([]json.Number, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		if value, ok := attr.([]json.Number); ok {
			return value, nil
		}
		return nil, ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return nil, ErrKeyNotExists
	}
	var value []json.Number
	err := json.Unmarshal(raw, &value)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return nil, ErrTypeNotMatch
		}
		return nil, err
	}
	return value, nil
}

func (obj *Object) GetNullInt(key string) (NullInt, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		if ni, ok := attr.(NullInt); ok {
			return ni, nil
		}

		if ni, ok := attr.(*NullInt); ok {
			return *ni, nil
		}

		v := reflect.ValueOf(attr)
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
			return NullInt{Null: false, Value: int(v.Int())}, nil
		}
		return NullInt{}, ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return NullInt{}, ErrKeyNotExists
	}
	var value NullInt
	err := json.Unmarshal(raw, &value)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return value, ErrTypeNotMatch
		}
		return value, err
	}
	return value, nil
}

func (obj *Object) GetBool(key string) (bool, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		if boolean, ok := attr.(bool); ok {
			return boolean, nil
		}
		return false, ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return false, ErrKeyNotExists
	}
	var boolean bool
	err := json.Unmarshal(raw, &boolean)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return false, ErrTypeNotMatch
		}
		return false, err
	}
	return boolean, nil
}

func (obj *Object) GetNullBool(key string) (NullBool, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		if ni, ok := attr.(bool); ok {
			return NullBool{Null: false, Value: ni}, nil
		}

		if ni, ok := attr.(NullBool); ok {
			return ni, nil
		}

		if ni, ok := attr.(*NullBool); ok {
			return *ni, nil
		}

		return NullBool{}, ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return NullBool{}, ErrKeyNotExists
	}
	var value NullBool
	err := json.Unmarshal(raw, &value)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return value, ErrTypeNotMatch
		}
		return value, err
	}
	return value, nil
}

func (obj *Object) GetObject(key string) (*Object, error) {
	raw, ok := obj.getRaw(key)
	if !ok {
		return nil, ErrKeyNotExists
	}
	var o Object
	err := json.Unmarshal(raw, &o)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return nil, ErrTypeNotMatch
		}
		return nil, err
	}
	// 引用类型的数据放入attrs中，如果引用数据更新
	// attrs会覆盖原有数据，实现更新
	obj.setAttr(key, &o)
	return &o, nil
}

func (obj *Object) GetAny(key string) (any, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		return attr, nil
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return nil, ErrKeyNotExists
	}
	var value any
	err := json.Unmarshal(raw, &value)
	if err != nil {
		return value, err
	}
	return value, nil
}

func (obj *Object) GetInts(key string) ([]int, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		if ints, ok := attr.([]int); ok {
			return ints, nil
		}
		return nil, ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return nil, ErrKeyNotExists
	}
	var ints []int
	err := json.Unmarshal(raw, &ints)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return nil, ErrTypeNotMatch
		}
		return nil, err
	}
	return ints, nil
}

func (obj *Object) GetBools(key string) ([]bool, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		if values, ok := attr.([]bool); ok {
			return values, nil
		}
		return nil, ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return nil, ErrKeyNotExists
	}
	var values []bool
	err := json.Unmarshal(raw, &values)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return nil, ErrTypeNotMatch
		}
		return nil, err
	}
	return values, nil
}

func (obj *Object) GetStrings(key string) ([]string, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		if values, ok := attr.([]string); ok {
			return values, nil
		}
		return nil, ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return nil, ErrKeyNotExists
	}
	var values []string
	err := json.Unmarshal(raw, &values)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return nil, ErrTypeNotMatch
		}
		return nil, err
	}
	return values, nil
}

func (obj *Object) GetObjects(key string) ([]*Object, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		if values, ok := attr.([]*Object); ok {
			return values, nil
		}
		return nil, ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return nil, ErrKeyNotExists
	}
	var values []*Object
	err := json.Unmarshal(raw, &values)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return nil, ErrTypeNotMatch
		}
		return nil, err
	}
	return values, nil
}

func (obj *Object) GetAnys(key string) ([]any, error) {
	attr, ok := obj.getAttr(key)
	if ok {
		if values, ok := attr.([]any); ok {
			return values, nil
		}
		return nil, ErrTypeNotMatch
	}

	raw, ok := obj.getRaw(key)
	if !ok {
		return nil, ErrKeyNotExists
	}
	var values []any
	err := json.Unmarshal(raw, &values)
	if err != nil {
		var e *json.UnmarshalTypeError
		if errors.As(err, &e) {
			return nil, ErrTypeNotMatch
		}
		return nil, err
	}
	return values, nil
}

func (obj *Object) Set(key string, value any) {
	obj.setAttr(key, value)
}

func (obj *Object) getRaw(key string) (json.RawMessage, bool) {
	if obj.raw == nil {
		return nil, false
	}
	raw, ok := obj.raw[key]
	return raw, ok
}

func (obj *Object) setRaw(key string, raw json.RawMessage) {
	if obj.raw == nil {
		obj.raw = make(map[string]json.RawMessage)
	}
	obj.locker.Lock()
	defer obj.locker.Unlock()
	obj.raw[key] = raw
}

func (obj *Object) getAttr(key string) (any, bool) {
	if obj.attrs == nil {
		return nil, false
	}
	raw, ok := obj.attrs[key]
	return raw, ok
}

func (obj *Object) setAttr(key string, value any) {
	if obj.attrs == nil {
		obj.attrs = make(map[string]any)
	}
	obj.locker.Lock()
	defer obj.locker.Unlock()
	obj.attrs[key] = value
}

func (obj *Object) UnmarshalJSON(b []byte) error {
	raw := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &raw)
	obj.raw = raw
	if obj.locker == nil {
		obj.locker = &sync.Mutex{}
	}
	return err
}

func (obj Object) MarshalJSON() ([]byte, error) {
	for k, attr := range obj.attrs {
		b, err := json.Marshal(attr)
		if err != nil {
			return nil, err
		}
		obj.setRaw(k, b)
	}
	return json.Marshal(obj.raw)
}

func NewObject() *Object {
	return &Object{
		locker: &sync.Mutex{},
		raw:    make(map[string]json.RawMessage),
		attrs:  make(map[string]any),
	}
}
