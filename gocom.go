package gocom

// Max 返回较大的值
func Max(x, y interface{}) interface{} {
	switch x.(type) {
	case int:
		if x.(int) > y.(int) {
			return x
		}
		return y
	case int32:
		if x.(int32) > y.(int32) {
			return x
		}
		return y
	case int64:
		if x.(int64) > y.(int64) {
			return x
		}
		return y
	case float32:
		if x.(float32) > y.(float32) {
			return x
		}
		return y
	case float64:
		if x.(float64) > y.(float64) {
			return x
		}
		return y
	default:
		panic("type error")
	}
}

// Min 返回较小的值
func Min(x, y interface{}) interface{} {
	switch x.(type) {
	case int:
		if x.(int) < y.(int) {
			return x
		}
		return y
	case int32:
		if x.(int32) < y.(int32) {
			return x
		}
		return y
	case int64:
		if x.(int64) < y.(int64) {
			return x
		}
		return y
	case float32:
		if x.(float32) < y.(float32) {
			return x
		}
		return y
	case float64:
		if x.(float64) < y.(float64) {
			return x
		}
		return y
	default:
		panic("type error")
	}
}
