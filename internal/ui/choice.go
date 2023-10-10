package ui

import "strconv"

type Choice struct {
	Group    string
	Label    string
	Value    string
	Selected bool
}

func (c *Choice) Val() string {
	return c.Value
}

func (c *Choice) L() string {
	return c.Label
}

func (c *Choice) IsSelected(x any) bool {
	switch x := x.(type) {
	case int64:
		return c.Value == strconv.FormatInt(x, 10)
	case int:
		return c.Value == strconv.Itoa(x)
	case string:
		return c.Value == x
	default:
		return false
	}
}
