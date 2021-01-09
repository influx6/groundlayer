package base

import (
	gocolorful "github.com/lucasb-eyer/go-colorful"
)

type ValuesDef struct {
	def    ScaleDef
	values map[interface{}]interface{}
}

func (sv *ValuesDef) GetScale() ScaleDef {
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

func (sv *ValuesDef) KeyAtIndex(index int) interface{} {
	if !sv.def.HasIndex(index) {
		panic("scale index out of bounds")
	}
	return sv.def.Index(index)
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
	Scale *ValuesDef
}

func NewString(scale ScaleDef) *String {
	return &String{
		Scale: NewValuesDef(scale),
	}
}

func (sv *String) GetScale() ScaleDef {
	return sv.Scale.GetScale()
}

// KeyAtIndex uses zero-index based values.
func (skv *String) KeyAtIndex(k int) string {
	return skv.Scale.KeyAtIndex(k).(string)
}

func (skv *String) ValueAtKey(k string) string {
	return skv.Scale.ValueAtKey(k).(string)
}

// ValueAtIndex uses zero-index based values.
func (skv *String) ValueAtIndex(i int) string {
	return skv.Scale.ValueAtIndex(i).(string)
}

// Index uses zero-index based values.
func (skv *String) Index(i int, v string) *String {
	skv.Scale.Index(i, v)
	return skv
}

func (skv *String) Key(k string, v string) *String {
	skv.Scale.Key(k, v)
	return skv
}

type StringNumber struct {
	Scale *ValuesDef
}

func NewStringNumber(scale ScaleDef) *StringNumber {
	return &StringNumber{
		Scale: NewValuesDef(scale),
	}
}

func (sv *StringNumber) GetScale() ScaleDef {
	return sv.Scale.GetScale()
}

func (skv *StringNumber) KeyAtIndex(k int) string {
	return skv.Scale.KeyAtIndex(k).(string)
}

func (skv *StringNumber) ValueAtKey(k string) int {
	return skv.Scale.ValueAtKey(k).(int)
}

func (skv *StringNumber) ValueAtIndex(i int) int {
	return skv.Scale.ValueAtIndex(i).(int)
}

func (skv *StringNumber) Index(i int, v int) *StringNumber {
	skv.Scale.Index(i, v)
	return skv
}

func (skv *StringNumber) Key(k string, v int) *StringNumber {
	skv.Scale.Key(k, v)
	return skv
}

type StringColor struct {
	Scale *ValuesDef
}

func NewStringColor(scale ScaleDef) *StringColor {
	return &StringColor{
		Scale: NewValuesDef(scale),
	}
}

func (skv *StringColor) ValueAtKey(k string) gocolorful.Color {
	return skv.Scale.ValueAtKey(k).(gocolorful.Color)
}

func (skv *StringColor) KeyAtIndex(k int) string {
	return skv.Scale.KeyAtIndex(k).(string)
}

func (skv *StringColor) ValueAtIndex(i int) gocolorful.Color {
	return skv.Scale.ValueAtIndex(i).(gocolorful.Color)
}

func (skv *StringColor) Index(i int, v gocolorful.Color) *StringColor {
	skv.Scale.Index(i, v)
	return skv
}

func (skv *StringColor) Key(k string, v gocolorful.Color) *StringColor {
	skv.Scale.Key(k, v)
	return skv
}

type StringPalette struct {
	Scale *ValuesDef
}

func NewStringPalette(scale ScaleDef) *StringPalette {
	return &StringPalette{
		Scale: NewValuesDef(scale),
	}
}

func (skv *StringPalette) KeyAtIndex(k int) string {
	return skv.Scale.KeyAtIndex(k).(string)
}

func (skv *StringPalette) ValueAtKey(k string) Palette {
	return skv.Scale.ValueAtKey(k).(Palette)
}

func (skv *StringPalette) ValueAtIndex(i int) Palette {
	return skv.Scale.ValueAtIndex(i).(Palette)
}

func (skv *StringPalette) Index(i int, v Palette) *StringPalette {
	skv.Scale.Index(i, v)
	return skv
}

func (skv *StringPalette) Key(k string, v Palette) *StringPalette {
	skv.Scale.Key(k, v)
	return skv
}

type StringFloat struct {
	Scale *ValuesDef
}

func NewStringFloat(scale ScaleDef) *StringFloat {
	return &StringFloat{
		Scale: NewValuesDef(scale),
	}
}

func (skv *StringFloat) KeyAtIndex(k int) string {
	return skv.Scale.KeyAtIndex(k).(string)
}

func (skv *StringFloat) ValueAtKey(k string) float64 {
	return skv.Scale.ValueAtKey(k).(float64)
}

func (skv *StringFloat) ValueAtIndex(i int) float64 {
	return skv.Scale.ValueAtIndex(i).(float64)
}

func (skv *StringFloat) Index(i int, v float64) *StringFloat {
	skv.Scale.Index(i, v)
	return skv
}

func (skv *StringFloat) Key(k string, v float64) *StringFloat {
	skv.Scale.Key(k, v)
	return skv
}

type Number struct {
	Scale *ValuesDef
}

func NewNumber(scale ScaleDef) *Number {
	return &Number{
		Scale: NewValuesDef(scale),
	}
}

func (skv *Number) ValueAtKey(k int) int {
	return skv.Scale.ValueAtKey(k).(int)
}

func (skv *Number) KeyAtIndex(k int) string {
	return skv.Scale.KeyAtIndex(k).(string)
}

func (skv *Number) ValueAtIndex(i int) int {
	return skv.Scale.ValueAtIndex(i).(int)
}

func (skv *Number) Index(i int, v int) *Number {
	skv.Scale.Index(i, v)
	return skv
}

func (skv *Number) Key(k int, v int) *Number {
	skv.Scale.Key(k, v)
	return skv
}

type NumberPalette struct {
	Scale *ValuesDef
}

func NewNumberPalette(scale ScaleDef) *NumberPalette {
	return &NumberPalette{
		Scale: NewValuesDef(scale),
	}
}

func (skv *NumberPalette) KeyAtIndex(k int) string {
	return skv.Scale.KeyAtIndex(k).(string)
}

func (skv *NumberPalette) ValueAtKey(k int) Palette {
	return skv.Scale.ValueAtKey(k).(Palette)
}

func (skv *NumberPalette) ValueAtIndex(i int) Palette {
	return skv.Scale.ValueAtIndex(i).(Palette)
}

func (skv *NumberPalette) Index(i int, v Palette) *NumberPalette {
	skv.Scale.Index(i, v)
	return skv
}

func (skv *NumberPalette) Key(k int, v Palette) *NumberPalette {
	skv.Scale.Key(k, v)
	return skv
}

type NumberColor struct {
	Scale *ValuesDef
}

func NewNumberColor(scale ScaleDef) *NumberColor {
	return &NumberColor{
		Scale: NewValuesDef(scale),
	}
}

func (skv *NumberColor) KeyAtIndex(k int) string {
	return skv.Scale.KeyAtIndex(k).(string)
}

func (skv *NumberColor) ValueAtKey(k int) gocolorful.Color {
	return skv.Scale.ValueAtKey(k).(gocolorful.Color)
}

func (skv *NumberColor) ValueAtIndex(i int) gocolorful.Color {
	return skv.Scale.ValueAtIndex(i).(gocolorful.Color)
}

func (skv *NumberColor) Index(i int, v gocolorful.Color) *NumberColor {
	skv.Scale.Index(i, v)
	return skv
}

func (skv *NumberColor) Key(k int, v gocolorful.Color) *NumberColor {
	skv.Scale.Key(k, v)
	return skv
}

type NumberFloat struct {
	Scale *ValuesDef
}

func NewNumberFloat(scale ScaleDef) *NumberFloat {
	return &NumberFloat{
		Scale: NewValuesDef(scale),
	}
}

func (skv *NumberFloat) KeyAtIndex(k int) string {
	return skv.Scale.KeyAtIndex(k).(string)
}

func (skv *NumberFloat) ValueAtKey(k int) float64 {
	return skv.Scale.ValueAtKey(k).(float64)
}

func (skv *NumberFloat) ValueAtIndex(i int) float64 {
	return skv.Scale.ValueAtIndex(i).(float64)
}

func (skv *NumberFloat) Index(i int, v float64) *NumberFloat {
	skv.Scale.Index(i, v)
	return skv
}

func (skv *NumberFloat) Key(k int, v float64) *NumberFloat {
	skv.Scale.Key(k, v)
	return skv
}

type NumberString struct {
	Scale *ValuesDef
}

func NewNumberString(scale ScaleDef) *NumberString {
	return &NumberString{
		Scale: NewValuesDef(scale),
	}
}

func (skv *NumberString) KeyAtIndex(k int) string {
	return skv.Scale.KeyAtIndex(k).(string)
}

func (skv *NumberString) ValueAtKey(k int) string {
	return skv.Scale.ValueAtKey(k).(string)
}

func (skv *NumberString) ValueAtIndex(i int) string {
	return skv.Scale.ValueAtIndex(i).(string)
}

func (skv *NumberString) Index(i int, v string) *NumberString {
	skv.Scale.Index(i, v)
	return skv
}

func (skv *NumberString) Key(k int, v string) *NumberString {
	skv.Scale.Key(k, v)
	return skv
}
