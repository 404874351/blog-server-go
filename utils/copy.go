package utils

import (
	"errors"
	"fmt"
	"reflect"
)

//
// CopyFields
//  @Description: 简单结构体属性浅拷贝 (&dst, src)，无法拷贝父级属性
//  @param dst
//  @param src
//  @return err
//
func CopyFields(dst interface{}, src interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}
	// 取指针的具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// src必须为结构体或者结构体指针
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// 复制结构体属性
	copyStructFields(dstType, dstValue, srcType, srcValue)
	return nil
}

func copyStructFields(dstType reflect.Type, dstValue reflect.Value, srcType reflect.Type, srcValue reflect.Value) {
	// 属性个数
	propertyNums := dstType.NumField()

	for i := 0; i < propertyNums; i++ {
		// 按索引获取结构体属性
		property := dstType.Field(i)
		// 待填充属性值
		propertyValue := srcValue.FieldByName(property.Name)

		// 无效，则src没有这个属性
		if !propertyValue.IsValid() {
			// 如果是父级继承，那么尝试递归复制
			if property.Type.Kind() == reflect.Struct && property.Anonymous {
				parent := dstValue.Field(i)
				copyStructFields(property.Type, parent, srcType, srcValue)
			}
			continue
		}
		// src有这个属性，但类型不同，也不能复制
		if property.Type != propertyValue.Type() {
			continue
		}
		// 可复制则直接复制
		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(propertyValue)
		}
	}
}
