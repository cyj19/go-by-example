/**
 * @Author: cyj19
 * @Date: 2021/11/15 18:27
 */

package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// 参数校验组件原理

type Nested struct {
	Email string `validate:"email"`
}

type T struct {
	Age    int `validate:"eq=10"`
	Nested Nested
}

func validateEmail(input string) bool {
	// \w\.]{2,10} : 数字或字母或下划线或句号至少出现2次不超过10次
	// \w+ : 数字或字母或下划线或句号出现1次或多次
	// [a-z]{2,4} : 字母a-z至少出现2次不超过4次
	pass, _ := regexp.MatchString(`^([\w.]{2,10})@(\w+).([a-z]{2,4})$`, input)
	return pass
}

func validate(v interface{}) (validateResult bool, errMsg string) {
	validateResult = true
	errMsg = "success"
	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	for i := 0; i < vv.NumField(); i++ {

		// 获取字段
		fieldVal := vv.Field(i)
		// 获取字段上的validate标签值
		tagContent := vt.Field(i).Tag.Get("validate")
		// 获取字段类型
		k := fieldVal.Kind()
		switch k {
		case reflect.Int:
			val := fieldVal.Int()
			tagValStr := strings.Split(tagContent, "=")
			tagVal, _ := strconv.ParseInt(tagValStr[1], 10, 64)
			if val != tagVal {
				errMsg = "validate int failed, tag is:" + strconv.FormatInt(tagVal, 10)
				validateResult = false
			}
		case reflect.String:
			val := fieldVal.String()
			tagValStr := tagContent
			switch tagValStr {
			case "email":
				validateResult = validateEmail(val)
				if !validateResult {
					errMsg = "validate email failed"
				}
			}
		case reflect.Struct:
			val := fieldVal.Interface()
			// 递归调用
			validateResult, errMsg = validate(val)

		}
		if !validateResult {
			break
		}
	}
	return
}

func main() {
	testData := T{Age: 10, Nested: Nested{
		Email: "abc123@abc.com",
	}}

	result, msg := validate(testData)
	fmt.Printf("result=%t msg=%s \n", result, msg)
}
