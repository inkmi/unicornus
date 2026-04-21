package uni

import "strconv"

// Choice is one selectable option in a select, radio, or multi-check field.
// Label is user-visible text; Value is the form-submit value. Group is the
// optional optgroup key. Checked reflects whether the choice is selected
// for the current value.
type Choice struct {
	Group   string
	Label   string
	Value   string
	Checked bool
}

// Val returns the form-submit Value for this Choice.
func (c *Choice) Val() string {
	return c.Value
}

// L returns the user-visible Label for this Choice.
func (c *Choice) L() string {
	return c.Label
}

// IsSelected reports whether the given value matches this Choice.
// Supported value kinds: int, int64, and string.
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
