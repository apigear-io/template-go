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
func AsSameStruct1Interface(v any) (SameStruct1Interface, error) {
    switch v := v.(type) {
    case SameStruct1Interface:
        return v, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to SameStruct1Interface", v, v)
    }
}

func AsSameStruct1InterfaceArray(v any) ([]SameStruct1Interface, error) {
    switch v := v.(type) {
    case []SameStruct1Interface:
        return v, nil
    case []interface{}:
        result := make([]SameStruct1Interface, len(v))
        for i, value := range v {
            result[i], _ = AsSameStruct1Interface(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []SameStruct1Interface", v, v)
    }
}
func AsSameStruct2Interface(v any) (SameStruct2Interface, error) {
    switch v := v.(type) {
    case SameStruct2Interface:
        return v, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to SameStruct2Interface", v, v)
    }
}

func AsSameStruct2InterfaceArray(v any) ([]SameStruct2Interface, error) {
    switch v := v.(type) {
    case []SameStruct2Interface:
        return v, nil
    case []interface{}:
        result := make([]SameStruct2Interface, len(v))
        for i, value := range v {
            result[i], _ = AsSameStruct2Interface(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []SameStruct2Interface", v, v)
    }
}
func AsSameEnum1Interface(v any) (SameEnum1Interface, error) {
    switch v := v.(type) {
    case SameEnum1Interface:
        return v, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to SameEnum1Interface", v, v)
    }
}

func AsSameEnum1InterfaceArray(v any) ([]SameEnum1Interface, error) {
    switch v := v.(type) {
    case []SameEnum1Interface:
        return v, nil
    case []interface{}:
        result := make([]SameEnum1Interface, len(v))
        for i, value := range v {
            result[i], _ = AsSameEnum1Interface(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []SameEnum1Interface", v, v)
    }
}
func AsSameEnum2Interface(v any) (SameEnum2Interface, error) {
    switch v := v.(type) {
    case SameEnum2Interface:
        return v, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to SameEnum2Interface", v, v)
    }
}

func AsSameEnum2InterfaceArray(v any) ([]SameEnum2Interface, error) {
    switch v := v.(type) {
    case []SameEnum2Interface:
        return v, nil
    case []interface{}:
        result := make([]SameEnum2Interface, len(v))
        for i, value := range v {
            result[i], _ = AsSameEnum2Interface(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []SameEnum2Interface", v, v)
    }
}
func AsStruct1(v any) (Struct1, error) {
    switch v := v.(type) {
    case Struct1:
        return v, nil
    default:
        return Struct1{}, fmt.Errorf("unable to cast %#v of type %T to Struct1", v, v)
    }
}

func AsStruct1Array(v any) ([]Struct1, error) {
    switch v := v.(type) {
    case []Struct1:
        return v, nil
    case []interface{}:
        result := make([]Struct1, len(v))
        for i, value := range v {
            result[i], _ = AsStruct1(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []Struct1", v, v)
    }
}
func AsStruct2(v any) (Struct2, error) {
    switch v := v.(type) {
    case Struct2:
        return v, nil
    default:
        return Struct2{}, fmt.Errorf("unable to cast %#v of type %T to Struct2", v, v)
    }
}

func AsStruct2Array(v any) ([]Struct2, error) {
    switch v := v.(type) {
    case []Struct2:
        return v, nil
    case []interface{}:
        result := make([]Struct2, len(v))
        for i, value := range v {
            result[i], _ = AsStruct2(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []Struct2", v, v)
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

