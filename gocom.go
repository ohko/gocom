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

// Ternary 三目运算
func Ternary(b bool, x, y interface{}) interface{} {
	if b {
		return x
	}
	return y
}

// Type 获取对象的类型
func Type(o interface{}) string {
	switch o.(type) {
	case int:
		return "int"
	case int8:
		return "int8"
	case int16:
		return "int16"
	case int32:
		return "int32"
	case int64:
		return "int64"
	case uint:
		return "uint"
	case uint8:
		return "uint8"
	case uint16:
		return "uint16"
	case uint32:
		return "uint32"
	case uint64:
		return "uint64"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "interface{}"
	}
}
