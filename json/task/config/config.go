package config

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*
	序列化和反序列化
*/

func MarshalFile(data interface{}) error {
	f, err := os.OpenFile("./config.ini", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	buf, err := Marshal(data)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	w.Write(buf)
	// 刷新缓存区，写入文件
	w.Flush()
	return nil
}

func UnMarshalFile(result interface{}) error {
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		return err
	}
	return UnMarshal(data, result)
}

/*
	序列化
	struct --> []byte
	基本思路：反射获取结构体的字段和标签并转为[]byte
*/
func Marshal(data interface{}) ([]byte, error) {

	// 反射获取data的类型和值
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	// 判断data的具体类型
	if t.Kind() != reflect.Struct {
		return nil, errors.New("data type is not a struct")
	}

	var conf []string
	// 获取data的字段
	for i := 0; i < t.NumField(); i++ {
		// 获取字段和值
		labelField := t.Field(i)
		labelVal := v.Field(i)
		fieldType := labelField.Type
		if fieldType.Kind() != reflect.Struct {
			continue
		}

		// 序列化Server、Mysql
		// 获取字段标签ini对应的值
		tagVal := labelField.Tag.Get("ini")
		// 如果无标签，默认和字段名称一致
		if len(tagVal) == 0 {
			tagVal = labelField.Name
		}
		// 拼接为[server]或[mysql]，保证是单独一行
		label := fmt.Sprintf("\n[%s]\n", tagVal)
		conf = append(conf, label)

		// 解析Server或Mysql
		for j := 0; j < fieldType.NumField(); j++ {
			// 获取字段和值
			keyField := fieldType.Field(j)
			valField := labelVal.Field(j)
			fieldTagVal := keyField.Tag.Get("ini")
			if len(fieldTagVal) == 0 {
				fieldTagVal = keyField.Name
			}
			// 拼接为xx = yy
			item := fmt.Sprintf("%s = %v\n", fieldTagVal, valField.Interface())
			conf = append(conf, item)
		}

	}

	// 转为[]byte
	var result []byte
	for _, value := range conf {
		result = append(result, []byte(value)...)
	}

	return result, nil

}

/*
	反序列化
	[]byte --> struct
	基本思路：[]byte转为string，根据\n进行切割，再按照标签赋值给字段
*/

func UnMarshal(data []byte, result interface{}) error {
	// 获取result类型
	t := reflect.TypeOf(result)
	// result必须是指针且指向的元素必须是结构体
	if t.Kind() != reflect.Ptr {
		return errors.New("result is not a pointer")
	}
	if t.Elem().Kind() != reflect.Struct {
		return errors.New("element is not a struct")
	}

	var fieldName string
	// 根据\n进行切割
	lineArr := strings.Split(string(data), "\n")
	for _, line := range lineArr {
		// 过滤前后空格
		line = strings.TrimSpace(line)
		// 过滤文档注释
		if len(line) == 0 || line[0] == '#' || line[0] == ';' {
			continue
		}
		// 处理头标签
		if line[0] == '[' {
			fieldName = doHeadTag(line, t.Elem())
			// 处理下一行
			continue
		}
		// 处理数据，赋值给结构体
		err := doField(fieldName, line, result)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
	处理头标签
	参数：行数据，字段类型
*/
func doHeadTag(line string, t reflect.Type) string {
	// 去除[]
	label := line[1 : len(line)-1]
	// 匹配字段
	var fieldName string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// 获取标签
		tag := field.Tag.Get("ini")
		if label == tag {
			fieldName = field.Name
			break
		}
	}
	return fieldName
}

/*
	处理行数据
	参数：头标签对应的字段名，行数据，结果指针
*/
func doField(fieldName string, line string, result interface{}) error {
	// t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)
	// 根据=切割
	key := strings.TrimSpace(line[0:strings.LastIndex(line, "=")])
	val := strings.TrimSpace(line[strings.LastIndex(line, "=")+1:])
	fmt.Println(key, val)
	// 根据字段名获取结构体值
	structVal := v.Elem().FieldByName(fieldName)
	// 获取结构体
	structType := structVal.Type()
	var itemFieldName string
	for i := 0; i < structType.NumField(); i++ {
		itemField := structType.Field(i)
		tag := itemField.Tag.Get("ini")
		if key == tag {
			itemFieldName = itemField.Name
			break
		}
	}
	// 获取结构体中的字段值
	itemFieldVal := structVal.FieldByName(itemFieldName)

	// 反射赋值
	switch itemFieldVal.Type().Kind() {
	case reflect.String:
		itemFieldVal.SetString(val)
	case reflect.Int:
		iVal, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		itemFieldVal.SetInt(int64(iVal))
	case reflect.Float64:
		fVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		itemFieldVal.SetFloat(fVal)
	}

	return nil
}
