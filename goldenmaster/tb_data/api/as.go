package api

import (
	"encoding/json"
	"fmt"
)


func AsInt(v any) (int64, error) {
	switch v := v.(type) {
	case int64:
		return v, nil
	case json.Number:
		return v.Int64()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", v, v)
	}
}

func AsIntArray(v any) ([]int64, error) {
	switch v := v.(type) {
	case []int64:
		return v, nil
	case []interface{}:
		result := make([]int64, len(v))
		for i, value := range v {
			result[i], _ = AsInt(value)
		}
		return result, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int64", v, v)
	}
}

func AsFloat(v any) (float64, error) {
	switch v := v.(type) {
	case float64:
		return v, nil
	case json.Number:
		return v.Float64()
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", v, v)
	}
}

func AsFloatArray(v any) ([]float64, error) {
	switch v := v.(type) {
	case []float64:
		return v, nil
	case []interface{}:
		result := make([]float64, len(v))
		for i, value := range v {
			result[i], _ = AsFloat(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []float64", v, v)
	}
}

func AsBool(v any) (bool, error) {
	switch v := v.(type) {
	case bool:
		return v, nil
	default:
		return false, fmt.Errorf("unable to cast %#v of type %T to bool", v, v)
	}
}

func AsBoolArray(v any) ([]bool, error) {
	switch v := v.(type) {
	case []bool:
		return v, nil
    case []interface{}:
        result := make([]bool, len(v))
        for i, value := range v {
            result[i], _ = AsBool(value)
        }
        return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []bool", v, v)
	}
}

func AsString(v any) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", v, v)
	}
}

func AsStringArray(v any) ([]string, error) {
	switch v := v.(type) {
	case []string:
		return v, nil
	case []interface{}:
		result := make([]string, len(v))
		for i, value := range v {
			result[i], _ = AsString(value)
		}
		return result, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []string", v, v)
	}
}
func AsStructInterface(v any) (StructInterface, error) {
    switch v := v.(type) {
    case StructInterface:
        return v, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to StructInterface", v, v)
    }
}

func AsStructInterfaceArray(v any) ([]StructInterface, error) {
    switch v := v.(type) {
    case []StructInterface:
        return v, nil
    case []interface{}:
        result := make([]StructInterface, len(v))
        for i, value := range v {
            result[i], _ = AsStructInterface(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []StructInterface", v, v)
    }
}
func AsStructArrayInterface(v any) (StructArrayInterface, error) {
    switch v := v.(type) {
    case StructArrayInterface:
        return v, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to StructArrayInterface", v, v)
    }
}

func AsStructArrayInterfaceArray(v any) ([]StructArrayInterface, error) {
    switch v := v.(type) {
    case []StructArrayInterface:
        return v, nil
    case []interface{}:
        result := make([]StructArrayInterface, len(v))
        for i, value := range v {
            result[i], _ = AsStructArrayInterface(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []StructArrayInterface", v, v)
    }
}
func AsStructBool(v any) (StructBool, error) {
    switch v := v.(type) {
    case StructBool:
        return v, nil
    default:
        return StructBool{}, fmt.Errorf("unable to cast %#v of type %T to StructBool", v, v)
    }
}

func AsStructBoolArray(v any) ([]StructBool, error) {
    switch v := v.(type) {
    case []StructBool:
        return v, nil
    case []interface{}:
        result := make([]StructBool, len(v))
        for i, value := range v {
            result[i], _ = AsStructBool(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []StructBool", v, v)
    }
}
func AsStructInt(v any) (StructInt, error) {
    switch v := v.(type) {
    case StructInt:
        return v, nil
    default:
        return StructInt{}, fmt.Errorf("unable to cast %#v of type %T to StructInt", v, v)
    }
}

func AsStructIntArray(v any) ([]StructInt, error) {
    switch v := v.(type) {
    case []StructInt:
        return v, nil
    case []interface{}:
        result := make([]StructInt, len(v))
        for i, value := range v {
            result[i], _ = AsStructInt(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []StructInt", v, v)
    }
}
func AsStructFloat(v any) (StructFloat, error) {
    switch v := v.(type) {
    case StructFloat:
        return v, nil
    default:
        return StructFloat{}, fmt.Errorf("unable to cast %#v of type %T to StructFloat", v, v)
    }
}

func AsStructFloatArray(v any) ([]StructFloat, error) {
    switch v := v.(type) {
    case []StructFloat:
        return v, nil
    case []interface{}:
        result := make([]StructFloat, len(v))
        for i, value := range v {
            result[i], _ = AsStructFloat(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []StructFloat", v, v)
    }
}
func AsStructString(v any) (StructString, error) {
    switch v := v.(type) {
    case StructString:
        return v, nil
    default:
        return StructString{}, fmt.Errorf("unable to cast %#v of type %T to StructString", v, v)
    }
}

func AsStructStringArray(v any) ([]StructString, error) {
    switch v := v.(type) {
    case []StructString:
        return v, nil
    case []interface{}:
        result := make([]StructString, len(v))
        for i, value := range v {
            result[i], _ = AsStructString(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []StructString", v, v)
    }
}

