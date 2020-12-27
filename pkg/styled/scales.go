package styled

import gocolorful "github.com/lucasb-eyer/go-colorful"

type ScaleDef interface {
	Count() int
	Index(int) interface{}
	HasIndex(int) bool
	Key(interface{}) bool
	IsKeyType(interface{}) bool
}

var _ ScaleDef = (*AnyScaleDef)(nil)

func AnyScale(values ...interface{}) *AnyScaleDef {
	var scale AnyScaleDef
	scale.count = len(values)
	scale.index = map[int]interface{}{}
	scale.keys = map[interface{}]bool{}
	for index, value := range values {
		scale.index[index] = value
		scale.keys[value] = true
	}
	return &scale
}

type AnyScaleDef struct {
	count int
	index map[int]interface{}
	keys  map[interface{}]bool
}

func (s *AnyScaleDef) Count() int {
	return s.count
}

func (s *AnyScaleDef) Index(i int) interface{} {
	return s.index[i]
}

func (s *AnyScaleDef) HasIndex(i int) bool {
	_, hasIndex := s.index[i]
	return hasIndex
}

func (s *AnyScaleDef) Key(i interface{}) bool {
	return s.keys[i]
}

// IsKeyType always returns true for AnyScaleDef.
func (s *AnyScaleDef) IsKeyType(i interface{}) bool {
	return true
}

var _ ScaleDef = (*NumberScaleDef)(nil)

func NumberScale(values ...int) *NumberScaleDef {
	var scale NumberScaleDef
	scale.count = len(values)
	scale.index = map[int]int{}
	scale.keys = map[int]bool{}
	for index, value := range values {
		scale.index[index] = value
		scale.keys[value] = true
	}
	return &scale
}

type NumberScaleDef struct {
	count int
	index map[int]int
	keys  map[int]bool
}

func (s *NumberScaleDef) Count() int {
	return s.count
}

func (s *NumberScaleDef) Index(i int) interface{} {
	return s.index[i]
}

func (s *NumberScaleDef) HasIndex(i int) bool {
	_, hasIndex := s.index[i]
	return hasIndex
}

func (s *NumberScaleDef) Key(i interface{}) bool {
	if key, isInt := i.(int); isInt {
		return s.keys[key]
	}
	return false
}

func (s *NumberScaleDef) IsKeyType(i interface{}) bool {
	if _, isInt := i.(int); isInt {
		return true
	}
	return false
}

func (s *NumberScaleDef) IsValueType(i interface{}) bool {
	if _, isString := i.(string); isString {
		return true
	}
	return false
}

var _ ScaleDef = (*StringScaleDef)(nil)

func StringScale(values ...string) *StringScaleDef {
	var scale StringScaleDef
	scale.count = len(values)
	scale.index = map[int]string{}
	scale.keys = map[string]bool{}
	for index, value := range values {
		scale.index[index] = value
		scale.keys[value] = true
	}
	return &scale
}

type StringScaleDef struct {
	count int
	index map[int]string
	keys  map[string]bool
}

func (s *StringScaleDef) Count() int {
	return s.count
}

func (s *StringScaleDef) Index(i int) interface{} {
	return s.index[i]
}

func (s *StringScaleDef) HasIndex(i int) bool {
	_, hasIndex := s.index[i]
	return hasIndex
}

func (s *StringScaleDef) Key(i interface{}) bool {
	if stringKey, isString := i.(string); isString {
		return s.keys[stringKey]
	}
	return false
}

func (s *StringScaleDef) IsKeyType(i interface{}) bool {
	if _, isString := i.(string); isString {
		return true
	}
	return false
}

func (s *StringScaleDef) IsValueType(i interface{}) bool {
	if _, isString := i.(string); isString {
		return true
	}
	return false
}

type ValuesDef struct {
	def    ScaleDef
	values map[interface{}]interface{}
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

type StringKV struct {
	Values *ValuesDef
}

func NewStringKV(scale ScaleDef) *StringKV {
	return &StringKV{
		Values: NewValuesDef(scale),
	}
}

func (skv *StringKV) ValuesAtKey(k string) string {
	return skv.Values.ValueAtKey(k).(string)
}

func (skv *StringKV) ValuesAtIndex(i int) string {
	return skv.Values.ValueAtIndex(i).(string)
}

func (skv *StringKV) Index(i int, v string) *StringKV {
	skv.Values.Index(i, v)
	return skv
}

func (skv *StringKV) Key(k string, v string) *StringKV {
	skv.Values.Key(k, v)
	return skv
}

type StringNumberKV struct {
	Values *ValuesDef
}

func NewStringNumberKV(scale ScaleDef) *StringNumberKV {
	return &StringNumberKV{
		Values: NewValuesDef(scale),
	}
}

func (skv *StringNumberKV) ValuesAtKey(k string) int {
	return skv.Values.ValueAtKey(k).(int)
}

func (skv *StringNumberKV) ValuesAtIndex(i int) int {
	return skv.Values.ValueAtIndex(i).(int)
}

func (skv *StringNumberKV) Index(i int, v int) *StringNumberKV {
	skv.Values.Index(i, v)
	return skv
}

func (skv *StringNumberKV) Key(k string, v int) *StringNumberKV {
	skv.Values.Key(k, v)
	return skv
}

type StringColorKV struct {
	Values *ValuesDef
}

func NewStringColorKV(scale ScaleDef) *StringColorKV {
	return &StringColorKV{
		Values: NewValuesDef(scale),
	}
}

func (skv *StringColorKV) ValuesAtKey(k string) gocolorful.Color {
	return skv.Values.ValueAtKey(k).(gocolorful.Color)
}

func (skv *StringColorKV) ValuesAtIndex(i int) gocolorful.Color {
	return skv.Values.ValueAtIndex(i).(gocolorful.Color)
}

func (skv *StringColorKV) Index(i int, v gocolorful.Color) *StringColorKV {
	skv.Values.Index(i, v)
	return skv
}

func (skv *StringColorKV) Key(k string, v gocolorful.Color) *StringColorKV {
	skv.Values.Key(k, v)
	return skv
}

type StringPaletteKV struct {
	Values *ValuesDef
}

func NewStringPaletteKV(scale ScaleDef) *StringPaletteKV {
	return &StringPaletteKV{
		Values: NewValuesDef(scale),
	}
}

func (skv *StringPaletteKV) ValuesAtKey(k string) Palette {
	return skv.Values.ValueAtKey(k).(Palette)
}

func (skv *StringPaletteKV) ValuesAtIndex(i int) Palette {
	return skv.Values.ValueAtIndex(i).(Palette)
}

func (skv *StringPaletteKV) Index(i int, v Palette) *StringPaletteKV {
	skv.Values.Index(i, v)
	return skv
}

func (skv *StringPaletteKV) Key(k string, v Palette) *StringPaletteKV {
	skv.Values.Key(k, v)
	return skv
}

type StringFloatKV struct {
	Values *ValuesDef
}

func NewStringFloatKV(scale ScaleDef) *StringFloatKV {
	return &StringFloatKV{
		Values: NewValuesDef(scale),
	}
}

func (skv *StringFloatKV) ValuesAtKey(k string) float64 {
	return skv.Values.ValueAtKey(k).(float64)
}

func (skv *StringFloatKV) ValuesAtIndex(i int) float64 {
	return skv.Values.ValueAtIndex(i).(float64)
}

func (skv *StringFloatKV) Index(i int, v float64) *StringFloatKV {
	skv.Values.Index(i, v)
	return skv
}

func (skv *StringFloatKV) Key(k string, v float64) *StringFloatKV {
	skv.Values.Key(k, v)
	return skv
}

type NumberKV struct {
	Values *ValuesDef
}

func NewNumberKV(scale ScaleDef) *NumberKV {
	return &NumberKV{
		Values: NewValuesDef(scale),
	}
}

func (skv *NumberKV) ValuesAtKey(k int) int {
	return skv.Values.ValueAtKey(k).(int)
}

func (skv *NumberKV) ValuesAtIndex(i int) int {
	return skv.Values.ValueAtIndex(i).(int)
}

func (skv *NumberKV) Index(i int, v int) *NumberKV {
	skv.Values.Index(i, v)
	return skv
}

func (skv *NumberKV) Key(k int, v int) *NumberKV {
	skv.Values.Key(k, v)
	return skv
}

type NumberPaletteKV struct {
	Values *ValuesDef
}

func NewNumberPaletteKV(scale ScaleDef) *NumberPaletteKV {
	return &NumberPaletteKV{
		Values: NewValuesDef(scale),
	}
}

func (skv *NumberPaletteKV) ValuesAtKey(k int) Palette {
	return skv.Values.ValueAtKey(k).(Palette)
}

func (skv *NumberPaletteKV) ValuesAtIndex(i int) Palette {
	return skv.Values.ValueAtIndex(i).(Palette)
}

func (skv *NumberPaletteKV) Index(i int, v Palette) *NumberPaletteKV {
	skv.Values.Index(i, v)
	return skv
}

func (skv *NumberPaletteKV) Key(k int, v Palette) *NumberPaletteKV {
	skv.Values.Key(k, v)
	return skv
}

type NumberColorKV struct {
	Values *ValuesDef
}

func NewNumberColorKV(scale ScaleDef) *NumberColorKV {
	return &NumberColorKV{
		Values: NewValuesDef(scale),
	}
}

func (skv *NumberColorKV) ValuesAtKey(k int) gocolorful.Color {
	return skv.Values.ValueAtKey(k).(gocolorful.Color)
}

func (skv *NumberColorKV) ValuesAtIndex(i int) gocolorful.Color {
	return skv.Values.ValueAtIndex(i).(gocolorful.Color)
}

func (skv *NumberColorKV) Index(i int, v gocolorful.Color) *NumberColorKV {
	skv.Values.Index(i, v)
	return skv
}

func (skv *NumberColorKV) Key(k int, v gocolorful.Color) *NumberColorKV {
	skv.Values.Key(k, v)
	return skv
}

type NumberFloatKV struct {
	Values *ValuesDef
}

func NewNumberFloatKV(scale ScaleDef) *NumberFloatKV {
	return &NumberFloatKV{
		Values: NewValuesDef(scale),
	}
}

func (skv *NumberFloatKV) ValuesAtKey(k int) float64 {
	return skv.Values.ValueAtKey(k).(float64)
}

func (skv *NumberFloatKV) ValuesAtIndex(i int) float64 {
	return skv.Values.ValueAtIndex(i).(float64)
}

func (skv *NumberFloatKV) Index(i int, v float64) *NumberFloatKV {
	skv.Values.Index(i, v)
	return skv
}

func (skv *NumberFloatKV) Key(k int, v float64) *NumberFloatKV {
	skv.Values.Key(k, v)
	return skv
}

type NumberStringKV struct {
	Values *ValuesDef
}

func NewNumberStringKV(scale ScaleDef) *NumberStringKV {
	return &NumberStringKV{
		Values: NewValuesDef(scale),
	}
}

func (skv *NumberStringKV) ValuesAtKey(k int) string {
	return skv.Values.ValueAtKey(k).(string)
}

func (skv *NumberStringKV) ValuesAtIndex(i int) string {
	return skv.Values.ValueAtIndex(i).(string)
}

func (skv *NumberStringKV) Index(i int, v string) *NumberStringKV {
	skv.Values.Index(i, v)
	return skv
}

func (skv *NumberStringKV) Key(k int, v string) *NumberStringKV {
	skv.Values.Key(k, v)
	return skv
}
