package base

import (
	gocolorful "github.com/lucasb-eyer/go-colorful"
)

type ValuesDef struct {
	def    ScaleDef
	values map[interface{}]interface{}
}

func (sv *ValuesDef) Scale() ScaleDef {
	return sv.def
}

func NewValuesDef(scale ScaleDef) *ValuesDef {
	return &ValuesDef{
		def:    scale,
		values: map[interface{}]interface{}{},
	}
}

func (sv *ValuesDef) Key(key interface{}, value interface{}) *ValuesDef {
	if !sv.def.IsKeyType(key) {
		panic("Scale definition does not support this type of key")
	}
	sv.values[key] = value
	return sv
}

func (sv *ValuesDef) Index(index int, value interface{}) *ValuesDef {
	if !sv.def.HasIndex(index) {
		panic("scale index out of bounds")
	}
	var key = sv.def.Index(index)
	sv.values[key] = value
	return sv
}

func (sv *ValuesDef) ValueAtIndex(index int) interface{} {
	if !sv.def.HasIndex(index) {
		panic("scale index out of bounds")
	}
	var key = sv.def.Index(index)
	return sv.values[key]
}

func (sv *ValuesDef) ValueAtKey(key interface{}) interface{} {
	if !sv.def.Key(key) {
		panic("scale key is incorrect")
	}
	return sv.values[key]
}

type String struct {
	Values *ValuesDef
}

func NewString(scale ScaleDef) *String {
	return &String{
		Values: NewValuesDef(scale),
	}
}

func (sv *String) Scale() ScaleDef {
	return sv.Values.Scale()
}

func (skv *String) ValuesAtKey(k string) string {
	return skv.Values.ValueAtKey(k).(string)
}

func (skv *String) ValuesAtIndex(i int) string {
	return skv.Values.ValueAtIndex(i).(string)
}

func (skv *String) Index(i int, v string) *String {
	skv.Values.Index(i, v)
	return skv
}

func (skv *String) Key(k string, v string) *String {
	skv.Values.Key(k, v)
	return skv
}

type StringNumber struct {
	Values *ValuesDef
}

func NewStringNumber(scale ScaleDef) *StringNumber {
	return &StringNumber{
		Values: NewValuesDef(scale),
	}
}

func (sv *StringNumber) Scale() ScaleDef {
	return sv.Values.Scale()
}

func (skv *StringNumber) ValuesAtKey(k string) int {
	return skv.Values.ValueAtKey(k).(int)
}

func (skv *StringNumber) ValuesAtIndex(i int) int {
	return skv.Values.ValueAtIndex(i).(int)
}

func (skv *StringNumber) Index(i int, v int) *StringNumber {
	skv.Values.Index(i, v)
	return skv
}

func (skv *StringNumber) Key(k string, v int) *StringNumber {
	skv.Values.Key(k, v)
	return skv
}

type StringColor struct {
	Values *ValuesDef
}

func NewStringColor(scale ScaleDef) *StringColor {
	return &StringColor{
		Values: NewValuesDef(scale),
	}
}

func (skv *StringColor) ValuesAtKey(k string) gocolorful.Color {
	return skv.Values.ValueAtKey(k).(gocolorful.Color)
}

func (skv *StringColor) ValuesAtIndex(i int) gocolorful.Color {
	return skv.Values.ValueAtIndex(i).(gocolorful.Color)
}

func (skv *StringColor) Index(i int, v gocolorful.Color) *StringColor {
	skv.Values.Index(i, v)
	return skv
}

func (skv *StringColor) Key(k string, v gocolorful.Color) *StringColor {
	skv.Values.Key(k, v)
	return skv
}

type StringPalette struct {
	Values *ValuesDef
}

func NewStringPalette(scale ScaleDef) *StringPalette {
	return &StringPalette{
		Values: NewValuesDef(scale),
	}
}

func (skv *StringPalette) ValuesAtKey(k string) Palette {
	return skv.Values.ValueAtKey(k).(Palette)
}

func (skv *StringPalette) ValuesAtIndex(i int) Palette {
	return skv.Values.ValueAtIndex(i).(Palette)
}

func (skv *StringPalette) Index(i int, v Palette) *StringPalette {
	skv.Values.Index(i, v)
	return skv
}

func (skv *StringPalette) Key(k string, v Palette) *StringPalette {
	skv.Values.Key(k, v)
	return skv
}

type StringFloat struct {
	Values *ValuesDef
}

func NewStringFloat(scale ScaleDef) *StringFloat {
	return &StringFloat{
		Values: NewValuesDef(scale),
	}
}

func (skv *StringFloat) ValuesAtKey(k string) float64 {
	return skv.Values.ValueAtKey(k).(float64)
}

func (skv *StringFloat) ValuesAtIndex(i int) float64 {
	return skv.Values.ValueAtIndex(i).(float64)
}

func (skv *StringFloat) Index(i int, v float64) *StringFloat {
	skv.Values.Index(i, v)
	return skv
}

func (skv *StringFloat) Key(k string, v float64) *StringFloat {
	skv.Values.Key(k, v)
	return skv
}

type Number struct {
	Values *ValuesDef
}

func NewNumber(scale ScaleDef) *Number {
	return &Number{
		Values: NewValuesDef(scale),
	}
}

func (skv *Number) ValuesAtKey(k int) int {
	return skv.Values.ValueAtKey(k).(int)
}

func (skv *Number) ValuesAtIndex(i int) int {
	return skv.Values.ValueAtIndex(i).(int)
}

func (skv *Number) Index(i int, v int) *Number {
	skv.Values.Index(i, v)
	return skv
}

func (skv *Number) Key(k int, v int) *Number {
	skv.Values.Key(k, v)
	return skv
}

type NumberPalette struct {
	Values *ValuesDef
}

func NewNumberPalette(scale ScaleDef) *NumberPalette {
	return &NumberPalette{
		Values: NewValuesDef(scale),
	}
}

func (skv *NumberPalette) ValuesAtKey(k int) Palette {
	return skv.Values.ValueAtKey(k).(Palette)
}

func (skv *NumberPalette) ValuesAtIndex(i int) Palette {
	return skv.Values.ValueAtIndex(i).(Palette)
}

func (skv *NumberPalette) Index(i int, v Palette) *NumberPalette {
	skv.Values.Index(i, v)
	return skv
}

func (skv *NumberPalette) Key(k int, v Palette) *NumberPalette {
	skv.Values.Key(k, v)
	return skv
}

type NumberColor struct {
	Values *ValuesDef
}

func NewNumberColor(scale ScaleDef) *NumberColor {
	return &NumberColor{
		Values: NewValuesDef(scale),
	}
}

func (skv *NumberColor) ValuesAtKey(k int) gocolorful.Color {
	return skv.Values.ValueAtKey(k).(gocolorful.Color)
}

func (skv *NumberColor) ValuesAtIndex(i int) gocolorful.Color {
	return skv.Values.ValueAtIndex(i).(gocolorful.Color)
}

func (skv *NumberColor) Index(i int, v gocolorful.Color) *NumberColor {
	skv.Values.Index(i, v)
	return skv
}

func (skv *NumberColor) Key(k int, v gocolorful.Color) *NumberColor {
	skv.Values.Key(k, v)
	return skv
}

type NumberFloat struct {
	Values *ValuesDef
}

func NewNumberFloat(scale ScaleDef) *NumberFloat {
	return &NumberFloat{
		Values: NewValuesDef(scale),
	}
}

func (skv *NumberFloat) ValuesAtKey(k int) float64 {
	return skv.Values.ValueAtKey(k).(float64)
}

func (skv *NumberFloat) ValuesAtIndex(i int) float64 {
	return skv.Values.ValueAtIndex(i).(float64)
}

func (skv *NumberFloat) Index(i int, v float64) *NumberFloat {
	skv.Values.Index(i, v)
	return skv
}

func (skv *NumberFloat) Key(k int, v float64) *NumberFloat {
	skv.Values.Key(k, v)
	return skv
}

type NumberString struct {
	Values *ValuesDef
}

func NewNumberString(scale ScaleDef) *NumberString {
	return &NumberString{
		Values: NewValuesDef(scale),
	}
}

func (skv *NumberString) ValuesAtKey(k int) string {
	return skv.Values.ValueAtKey(k).(string)
}

func (skv *NumberString) ValuesAtIndex(i int) string {
	return skv.Values.ValueAtIndex(i).(string)
}

func (skv *NumberString) Index(i int, v string) *NumberString {
	skv.Values.Index(i, v)
	return skv
}

func (skv *NumberString) Key(k int, v string) *NumberString {
	skv.Values.Key(k, v)
	return skv
}
