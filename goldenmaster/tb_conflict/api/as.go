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
func AsConflict1(v any) (Conflict1, error) {
	switch v := v.(type) {
	case Conflict1:
		return v, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to Conflict1", v, v)
	}
}

func AsConflict1Array(v any) ([]Conflict1, error) {
	switch v := v.(type) {
	case []Conflict1:
		return v, nil
	case []interface{}:
		result := make([]Conflict1, len(v))
		for i, value := range v {
			result[i], _ = AsConflict1(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Conflict1", v, v)
	}
}
func AsConflict2(v any) (Conflict2, error) {
	switch v := v.(type) {
	case Conflict2:
		return v, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to Conflict2", v, v)
	}
}

func AsConflict2Array(v any) ([]Conflict2, error) {
	switch v := v.(type) {
	case []Conflict2:
		return v, nil
	case []interface{}:
		result := make([]Conflict2, len(v))
		for i, value := range v {
			result[i], _ = AsConflict2(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Conflict2", v, v)
	}
}
func AsConflict3(v any) (Conflict3, error) {
	switch v := v.(type) {
	case Conflict3:
		return v, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to Conflict3", v, v)
	}
}

func AsConflict3Array(v any) ([]Conflict3, error) {
	switch v := v.(type) {
	case []Conflict3:
		return v, nil
	case []interface{}:
		result := make([]Conflict3, len(v))
		for i, value := range v {
			result[i], _ = AsConflict3(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Conflict3", v, v)
	}
}
func AsConflict4(v any) (Conflict4, error) {
	switch v := v.(type) {
	case Conflict4:
		return v, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to Conflict4", v, v)
	}
}

func AsConflict4Array(v any) ([]Conflict4, error) {
	switch v := v.(type) {
	case []Conflict4:
		return v, nil
	case []interface{}:
		result := make([]Conflict4, len(v))
		for i, value := range v {
			result[i], _ = AsConflict4(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Conflict4", v, v)
	}
}
func AsConflict5(v any) (Conflict5, error) {
	switch v := v.(type) {
	case Conflict5:
		return v, nil
	default:
		return Conflict5{}, fmt.Errorf("unable to cast %#v of type %T to Conflict5", v, v)
	}
}

func AsConflict5Array(v any) ([]Conflict5, error) {
	switch v := v.(type) {
	case []Conflict5:
		return v, nil
	case []interface{}:
		result := make([]Conflict5, len(v))
		for i, value := range v {
			result[i], _ = AsConflict5(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Conflict5", v, v)
	}
}
func AsConflict6(v any) (Conflict6, error) {
	switch v := v.(type) {
	case Conflict6:
		return v, nil
	default:
		return Conflict6{}, fmt.Errorf("unable to cast %#v of type %T to Conflict6", v, v)
	}
}

func AsConflict6Array(v any) ([]Conflict6, error) {
	switch v := v.(type) {
	case []Conflict6:
		return v, nil
	case []interface{}:
		result := make([]Conflict6, len(v))
		for i, value := range v {
			result[i], _ = AsConflict6(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Conflict6", v, v)
	}
}
func AsConflict7(v any) (Conflict7, error) {
	switch v := v.(type) {
	case Conflict7:
		return v, nil
	default:
		return Conflict7{}, fmt.Errorf("unable to cast %#v of type %T to Conflict7", v, v)
	}
}

func AsConflict7Array(v any) ([]Conflict7, error) {
	switch v := v.(type) {
	case []Conflict7:
		return v, nil
	case []interface{}:
		result := make([]Conflict7, len(v))
		for i, value := range v {
			result[i], _ = AsConflict7(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Conflict7", v, v)
	}
}
func AsConflict8(v any) (Conflict8, error) {
	switch v := v.(type) {
	case Conflict8:
		return v, nil
	default:
		return Conflict8_UNKNOWN, fmt.Errorf("unable to cast %#v of type %T to Conflict8", v, v)
	}
}

func AsConflict8Array(v any) ([]Conflict8, error) {
	switch v := v.(type) {
	case []Conflict8:
		return v, nil
	case []interface{}:
		result := make([]Conflict8, len(v))
		for i, value := range v {
			result[i], _ = AsConflict8(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Conflict8", v, v)
	}
}
func AsConflict9(v any) (Conflict9, error) {
	switch v := v.(type) {
	case Conflict9:
		return v, nil
	default:
		return Conflict9_UNKNOWN, fmt.Errorf("unable to cast %#v of type %T to Conflict9", v, v)
	}
}

func AsConflict9Array(v any) ([]Conflict9, error) {
	switch v := v.(type) {
	case []Conflict9:
		return v, nil
	case []interface{}:
		result := make([]Conflict9, len(v))
		for i, value := range v {
			result[i], _ = AsConflict9(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Conflict9", v, v)
	}
}
func AsConflict10(v any) (Conflict10, error) {
	switch v := v.(type) {
	case Conflict10:
		return v, nil
	default:
		return Conflict10_UNKNOWN, fmt.Errorf("unable to cast %#v of type %T to Conflict10", v, v)
	}
}

func AsConflict10Array(v any) ([]Conflict10, error) {
	switch v := v.(type) {
	case []Conflict10:
		return v, nil
	case []interface{}:
		result := make([]Conflict10, len(v))
		for i, value := range v {
			result[i], _ = AsConflict10(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Conflict10", v, v)
	}
}
