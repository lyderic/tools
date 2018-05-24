package tools

/* poor version of C's ternary operator: condition ? a : b */
func ternary(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
