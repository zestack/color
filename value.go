package color

import "fmt"

type Value struct {
	value  any
	colors []Attribute
}

func NewValue(value any, attrs ...Attribute) *Value {
	return &Value{value, attrs}
}

func (c *Value) Add(attrs ...Attribute) *Value {
	c.colors = append(c.colors, attrs...)
	return c
}

func (c *Value) Format(f fmt.State, verb rune) {
	_, _ = f.Write(Bytes(c.colors...))
	s := fmt.Sprintf(fmt.FormatString(f, verb), c.value)
	_, _ = f.Write([]byte(s))
	_, _ = f.Write(resetBytes)
}

func (c *Value) Value() any {
	return c.value
}

func (c *Value) Bytes() []byte {
	buf := Bytes(c.colors...)
	buf = append(buf, fmt.Sprintf("%v", c.value)...)
	buf = append(buf, resetBytes...)
	return buf
}
