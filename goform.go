/**
 * 
 * @author wangchen
 * @version 2019-07-15 01:28
 */
package goform

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func MarshalForm(name string, desc string, submit string, schema reflect.Type) (interface{}, error) {
	schemaMap := make(map[string]interface{}, 0)
	for fIndex := 0; fIndex < schema.NumField(); fIndex++ {
		field := schema.Field(fIndex)
		tmp := make(map[string]interface{}, 0)
		tmp["type"] = field.Tag.Get("type")
		tmp["title"] = field.Tag.Get("title")
		if field.Tag.Get("enum") != "" {
			enums := strings.Split(field.Tag.Get("enum"), ",")
			tmp["enum"] = enums
		}

		fieldName := getJsonName(field)
		schemaMap[fieldName] = tmp
	}
	resultMap := map[string]interface{}{"name": name, "desc": desc, "submit": submit, "schema": schemaMap}
	return resultMap, nil
}

func UnmashalForm(form url.Values, obj interface{}) error {
	schema := reflect.ValueOf(obj).Elem().Type()
	for fIndex := 0; fIndex < schema.NumField(); fIndex++ {
		field := schema.Field(fIndex)
		fieldName := getJsonName(field)
		reflect.ValueOf(obj).Elem().FieldByName(field.Name).Set(reflect.ValueOf(form.Get(fieldName)))
	}
	return nil
}

func MarshalResponse(obj interface{}) (interface{}, error) {
	resultMap := make([]map[string]interface{}, 0)
	val := reflect.ValueOf(obj)
	schema := reflect.TypeOf(obj)
	for fIndex := 0; fIndex < schema.NumField(); fIndex++ {
		field := schema.Field(fIndex)
		tmp := map[string]interface{}{}
		switch field.Tag.Get("type") {
		case "text":
			tmp["type"] = "text"
			tmp["title"] = field.Tag.Get("title")
			tmp["data"] = fmt.Sprint(val.FieldByName(field.Name).Interface())
		case "json":
			tmp["type"] = "json"
			tmp["title"] = field.Tag.Get("title")
			tmp["data"] = val.FieldByName(field.Name).Interface()
		case "table":
			tmp["type"] = "table"
			tmp["title"] = field.Tag.Get("title")
			columnNames := make([]map[string]string, 0)
			tableDatas := val.FieldByName(field.Name)
			data := make(map[string]interface{}, 0)
			if tableDatas.Len() > 0 {
				columnTmp, err := toJsonMap(tableDatas.Index(0).Interface())
				if err != nil {
					return "", err
				}
				for key := range columnTmp {
					columnNames = append(columnNames, map[string]string{
						"title": key,
						"key":   key,
					})
				}
				data["columnNames"] = columnNames
				data["tableDatas"] = tableDatas.Interface()
			}
			tmp["data"] = data
		}
		resultMap = append(resultMap, tmp)
	}
	return resultMap, nil
}

func toJsonMap(obj interface{}) (map[string]interface{}, error) {
	jsonVal, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	columnTmp := make(map[string]interface{})
	json.Unmarshal(jsonVal, &columnTmp)
	return columnTmp, nil
}

func getJsonName(field reflect.StructField) string {
	fieldName := field.Name
	if jsonTag := field.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
		if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
			fieldName = jsonTag[:commaIdx]
		} else {
			fieldName = jsonTag
		}
	}
	return fieldName
}
