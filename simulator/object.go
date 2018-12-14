package simulator

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type Key struct {
	Namespace string
	Type      string
	Name      string
}

type Object struct {
	Key
	Value interface{}
}

func (o *Object) Struct(typ interface{}) error {
	if reflect.ValueOf(typ).Kind() != reflect.Ptr {
		return fmt.Errorf("typ must be pointer")
	}

	if _, ok := o.Value.(map[string]interface{}); ok {
		decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			TagName: "json",
			Result:  typ,
		})
		if err != nil {
			return err
		}

		err = decoder.Decode(o.Value)
		if err != nil {
			return err
		}

		o.Value = reflect.ValueOf(typ).Elem().Interface()
	} else {
		// TODO: Test ME!
		valueType := reflect.TypeOf(o.Value)
		typElem := reflect.ValueOf(typ).Elem()
		typType := typElem.Type()
		if valueType != typType {
			return fmt.Errorf("object value(%v:%v) and typ(%v:%v) have diffrent types",
				o.Value, valueType,
				typElem, typType)
		}
	}
	return nil
}

func (o *Object) Map() *Object {
	if _, ok := o.Value.(map[string]interface{}); !ok {
		var m map[string]interface{}
		if err := mapstructure.Decode(o.Value, &m); err == nil {
			o.Value = m
		}
	}
	return o
}
