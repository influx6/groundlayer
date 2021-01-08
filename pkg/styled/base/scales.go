package base

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
