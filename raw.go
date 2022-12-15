package jsontype

import "encoding/json"

func (obj *Object) ParseRaw(key string, v any) error {
	raw, ok := obj.getRaw(key)
	if !ok {
		return nil
	}
	return json.Unmarshal(raw, v)
}
