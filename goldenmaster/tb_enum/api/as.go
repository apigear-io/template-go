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
func AsEnumInterface(v any) (EnumInterface, error) {
	switch v := v.(type) {
	case EnumInterface:
		return v, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to EnumInterface", v, v)
	}
}

func AsEnumInterfaceArray(v any) ([]EnumInterface, error) {
	switch v := v.(type) {
	case []EnumInterface:
		return v, nil
	case []interface{}:
		result := make([]EnumInterface, len(v))
		for i, value := range v {
			result[i], _ = AsEnumInterface(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []EnumInterface", v, v)
	}
}
func AsEnum0(v any) (Enum0, error) {
	switch v := v.(type) {
	case Enum0:
		return v, nil
	default:
		return Enum0_UNKNOWN, fmt.Errorf("unable to cast %#v of type %T to Enum0", v, v)
	}
}

func AsEnum0Array(v any) ([]Enum0, error) {
	switch v := v.(type) {
	case []Enum0:
		return v, nil
	case []interface{}:
		result := make([]Enum0, len(v))
		for i, value := range v {
			result[i], _ = AsEnum0(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Enum0", v, v)
	}
}
func AsEnum1(v any) (Enum1, error) {
	switch v := v.(type) {
	case Enum1:
		return v, nil
	default:
		return Enum1_UNKNOWN, fmt.Errorf("unable to cast %#v of type %T to Enum1", v, v)
	}
}

func AsEnum1Array(v any) ([]Enum1, error) {
	switch v := v.(type) {
	case []Enum1:
		return v, nil
	case []interface{}:
		result := make([]Enum1, len(v))
		for i, value := range v {
			result[i], _ = AsEnum1(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Enum1", v, v)
	}
}
func AsEnum2(v any) (Enum2, error) {
	switch v := v.(type) {
	case Enum2:
		return v, nil
	default:
		return Enum2_UNKNOWN, fmt.Errorf("unable to cast %#v of type %T to Enum2", v, v)
	}
}

func AsEnum2Array(v any) ([]Enum2, error) {
	switch v := v.(type) {
	case []Enum2:
		return v, nil
	case []interface{}:
		result := make([]Enum2, len(v))
		for i, value := range v {
			result[i], _ = AsEnum2(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Enum2", v, v)
	}
}
func AsEnum3(v any) (Enum3, error) {
	switch v := v.(type) {
	case Enum3:
		return v, nil
	default:
		return Enum3_UNKNOWN, fmt.Errorf("unable to cast %#v of type %T to Enum3", v, v)
	}
}

func AsEnum3Array(v any) ([]Enum3, error) {
	switch v := v.(type) {
	case []Enum3:
		return v, nil
	case []interface{}:
		result := make([]Enum3, len(v))
		for i, value := range v {
			result[i], _ = AsEnum3(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Enum3", v, v)
	}
}
