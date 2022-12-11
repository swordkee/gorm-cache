package util

// If
//
//	@Description: if实现的三元表达式
//	@param boolExpression: 布尔表达式，最终返回一个布尔值
//	@param trueReturn: 当boolExpression返回值为true的时候返回的值
//	@param falseReturn: 当boolExpression返回值为false的时候返回的值
//	@return bool: 三元表达式的结果，为trueReturn或者falseReturn中的一个
func If[T any](boolExpression bool, trueReturn, falseReturn T) T {
	if boolExpression {
		return trueReturn
	} else {
		return falseReturn
	}
}

func In[E comparable](v E, s []E) bool {
	if Index(v, s) != -1 {
		return true
	}
	return false
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[E comparable](v E, s []E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}
