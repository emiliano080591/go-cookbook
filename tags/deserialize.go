package tags

import (
	"errors"
	"reflect"
	"strings"
)

func DeSerializeStructStrings(s string,res interface{}) error {
	r:=reflect.TypeOf(res)

	if r.Kind()!=reflect.Ptr {
		return errors.New("res must be a pointer")
	}
	r=r.Elem()
	value:=reflect.ValueOf(res).Elem()

	vals:=strings.Split(s,";")
	valMap:=make(map[string]string)
	for _,v:=range vals{
		keyval:=strings.Split(v,":")
		if len(keyval) !=2 {
			continue
		}
		valMap[keyval[0]]=keyval[1]
	}

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		// check if in the serialize set
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// ignore "-" otherwise that whole value
			// becomes the serialize 'key'
			if serialize == "-" {
				continue
			}
			// is it in the map
			if val, ok := valMap[serialize]; ok {
				value.Field(i).SetString(val)
			}
		} else if val, ok := valMap[field.Name]; ok {
			// is our field name in the map instead?
			value.Field(i).SetString(val)
		}
	}
	return nil
}