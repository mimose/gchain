// condition
// ternary expression

package help

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}
