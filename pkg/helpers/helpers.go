// Package helpers implements a series of helper functions which
// are only useful for go basic types, i.e strings, float, byte, rune, etc

package helpers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/influx6/groundlayer/pkg/domu"
)

func And(first, second interface{}) interface{} {
	var firstReflect = reflect.ValueOf(first)
	var secondReflect = reflect.ValueOf(second)
	return and(firstReflect, secondReflect)
}

func Or(first, second interface{}) interface{} {
	var firstReflect = reflect.ValueOf(first)
	var secondReflect = reflect.ValueOf(second)
	return and(firstReflect, secondReflect)
}

func Not(first interface{}) bool {
	var firstReflect = reflect.ValueOf(first)
	return not(firstReflect)
}

func Index(first interface{}, indexes ...int) (interface{}, error) {
	var firstReflect = reflect.ValueOf(first)
	var indexValues []reflect.Value
	for _, sec := range indexes {
		indexValues = append(indexValues, reflect.ValueOf(sec))
	}
	var res, err = index(firstReflect, indexValues...)
	if err != nil {
		return nil, err
	}
	return res.Elem(), nil
}

func Slice(first interface{}, indexes ...int) (interface{}, error) {
	var firstReflect = reflect.ValueOf(first)
	var indexValues []reflect.Value
	for _, sec := range indexes {
		indexValues = append(indexValues, reflect.ValueOf(sec))
	}
	var res, err = slice(firstReflect, indexValues...)
	if err != nil {
		return nil, err
	}
	return res.Elem(), nil
}

func Len(first interface{}) int {
	var firstReflect = reflect.ValueOf(first)
	var size, err = length(firstReflect)
	if err != nil {
		panic(err)
	}
	return size
}

func Cap(first, second interface{}) int {
	var firstReflect = reflect.ValueOf(first)
	return firstReflect.Cap()
}

func CanBeNill(first interface{}) bool {
	var firstReflect = reflect.TypeOf(first)
	return canBeNil(firstReflect)
}

func IsNill(first interface{}) bool {
	var firstReflect = reflect.ValueOf(first)
	return firstReflect.IsNil()
}

func IsInvalid(first, second interface{}) bool {
	var firstReflect = reflect.ValueOf(first)
	return firstReflect.IsNil() || !firstReflect.IsValid()
}

func Print(first interface{}) string {
	var firstReflect = reflect.ValueOf(first)
	return printValue(firstReflect)
}

func Printf(first interface{}, format string) string {
	var firstReflect = reflect.ValueOf(first)
	return printValueWithFormat(format, firstReflect)
}

func Println(first interface{}) string {
	var firstReflect = reflect.ValueOf(first)
	var content = printValue(firstReflect)
	return content + "\n"
}

func printValueWithFormat(format string, v reflect.Value) string {
	iface, ok := printableValue(v)
	if !ok {
		panic(fmt.Sprintf("can't print %s", v.Type()))
	}
	var buf strings.Builder
	_, err := fmt.Fprintf(&buf, format, iface)
	if err != nil {
		panic(err.Error())
	}
	return buf.String()
}

func printValue(v reflect.Value) string {
	iface, ok := printableValue(v)
	if !ok {
		panic(fmt.Sprintf("can't print %s", v.Type()))
	}
	var buf strings.Builder
	_, err := fmt.Fprint(&buf, iface)
	if err != nil {
		panic(err.Error())
	}
	return buf.String()
}

func URLQuery(first ...interface{}) string {
	return URLQueryEscaper(first...)
}

func AttachToNode(first interface{}, parent *domu.Node) {
	if first == nil {
		return
	}

	switch targetItem := first.(type) {
	case *domu.IntAttr:
		parent.Attrs.Add(targetItem)
	case *domu.StringAttr:
		parent.Attrs.Add(targetItem)
	case *domu.StringListAttr:
		parent.Attrs.Add(targetItem)
	case *domu.Node:
		parent.AppendChild(targetItem)
	case byte:
		parent.AppendChild(domu.Text(fmt.Sprintf("%+d", targetItem)))
	case complex128, complex64:
		parent.AppendChild(domu.Text(fmt.Sprintf("%+s", targetItem)))
	case string:
		parent.AppendChild(domu.Text(targetItem))
	case float32, float64:
		parent.AppendChild(domu.Text(fmt.Sprintf("%.4f", targetItem)))
	case uint, uint16, uint32, uint64:
		parent.AppendChild(domu.Text(fmt.Sprintf("%+d", targetItem)))
	case int, int16, int32, int64:
		parent.AppendChild(domu.Text(fmt.Sprintf("%+d", targetItem)))
	default:
		parent.AppendChild(domu.Text(fmt.Sprintf("%+s", targetItem)))
	}
}

func Equal(first, second interface{}) bool {
	var firstReflect = reflect.ValueOf(first)
	var secondReflect = reflect.ValueOf(second)
	var isTrue, err = eq(firstReflect, secondReflect)
	if err != nil {
		panic(err)
	}
	return isTrue
}

func NotEqual(first, second interface{}) bool {
	var firstReflect = reflect.ValueOf(first)
	var secondReflect = reflect.ValueOf(second)
	var isTrue, err = eq(firstReflect, secondReflect)
	if err != nil {
		panic(err)
	}
	return !isTrue
}

func LessThanEqualTo(first, second interface{}) bool {
	var firstReflect = reflect.ValueOf(first)
	var secondReflect = reflect.ValueOf(second)
	var isTrue, err = le(firstReflect, secondReflect)
	if err != nil {
		panic(err)
	}
	return isTrue
}

func LessThan(first, second interface{}) bool {
	var firstReflect = reflect.ValueOf(first)
	var secondReflect = reflect.ValueOf(second)
	var isTrue, err = lt(firstReflect, secondReflect)
	if err != nil {
		panic(err)
	}
	return isTrue
}

func GreaterThanEqualTo(first, second interface{}) bool {
	var firstReflect = reflect.ValueOf(first)
	var secondReflect = reflect.ValueOf(second)
	var isTrue, err = ge(firstReflect, secondReflect)
	if err != nil {
		panic(err)
	}
	return isTrue
}

func GreaterThan(first, second interface{}) bool {
	var firstReflect = reflect.ValueOf(first)
	var secondReflect = reflect.ValueOf(second)
	var isTrue, err = gt(firstReflect, secondReflect)
	if err != nil {
		panic(err)
	}
	return isTrue
}
