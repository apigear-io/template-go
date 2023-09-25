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
func AsManyParamInterface(v any) (ManyParamInterface, error) {
	switch v := v.(type) {
	case ManyParamInterface:
		return v, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to ManyParamInterface", v, v)
	}
}

func AsManyParamInterfaceArray(v any) ([]ManyParamInterface, error) {
	switch v := v.(type) {
	case []ManyParamInterface:
		return v, nil
	case []interface{}:
		result := make([]ManyParamInterface, len(v))
		for i, value := range v {
			result[i], _ = AsManyParamInterface(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []ManyParamInterface", v, v)
	}
}
func AsNestedStruct1Interface(v any) (NestedStruct1Interface, error) {
	switch v := v.(type) {
	case NestedStruct1Interface:
		return v, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to NestedStruct1Interface", v, v)
	}
}

func AsNestedStruct1InterfaceArray(v any) ([]NestedStruct1Interface, error) {
	switch v := v.(type) {
	case []NestedStruct1Interface:
		return v, nil
	case []interface{}:
		result := make([]NestedStruct1Interface, len(v))
		for i, value := range v {
			result[i], _ = AsNestedStruct1Interface(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []NestedStruct1Interface", v, v)
	}
}
func AsNestedStruct2Interface(v any) (NestedStruct2Interface, error) {
	switch v := v.(type) {
	case NestedStruct2Interface:
		return v, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to NestedStruct2Interface", v, v)
	}
}

func AsNestedStruct2InterfaceArray(v any) ([]NestedStruct2Interface, error) {
	switch v := v.(type) {
	case []NestedStruct2Interface:
		return v, nil
	case []interface{}:
		result := make([]NestedStruct2Interface, len(v))
		for i, value := range v {
			result[i], _ = AsNestedStruct2Interface(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []NestedStruct2Interface", v, v)
	}
}
func AsNestedStruct3Interface(v any) (NestedStruct3Interface, error) {
	switch v := v.(type) {
	case NestedStruct3Interface:
		return v, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to NestedStruct3Interface", v, v)
	}
}

func AsNestedStruct3InterfaceArray(v any) ([]NestedStruct3Interface, error) {
	switch v := v.(type) {
	case []NestedStruct3Interface:
		return v, nil
	case []interface{}:
		result := make([]NestedStruct3Interface, len(v))
		for i, value := range v {
			result[i], _ = AsNestedStruct3Interface(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []NestedStruct3Interface", v, v)
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
func AsStruct3(v any) (Struct3, error) {
	switch v := v.(type) {
	case Struct3:
		return v, nil
	default:
		return Struct3{}, fmt.Errorf("unable to cast %#v of type %T to Struct3", v, v)
	}
}

func AsStruct3Array(v any) ([]Struct3, error) {
	switch v := v.(type) {
	case []Struct3:
		return v, nil
	case []interface{}:
		result := make([]Struct3, len(v))
		for i, value := range v {
			result[i], _ = AsStruct3(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Struct3", v, v)
	}
}
func AsStruct4(v any) (Struct4, error) {
	switch v := v.(type) {
	case Struct4:
		return v, nil
	default:
		return Struct4{}, fmt.Errorf("unable to cast %#v of type %T to Struct4", v, v)
	}
}

func AsStruct4Array(v any) ([]Struct4, error) {
	switch v := v.(type) {
	case []Struct4:
		return v, nil
	case []interface{}:
		result := make([]Struct4, len(v))
		for i, value := range v {
			result[i], _ = AsStruct4(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []Struct4", v, v)
	}
}
func AsNestedStruct1(v any) (NestedStruct1, error) {
	switch v := v.(type) {
	case NestedStruct1:
		return v, nil
	default:
		return NestedStruct1{}, fmt.Errorf("unable to cast %#v of type %T to NestedStruct1", v, v)
	}
}

func AsNestedStruct1Array(v any) ([]NestedStruct1, error) {
	switch v := v.(type) {
	case []NestedStruct1:
		return v, nil
	case []interface{}:
		result := make([]NestedStruct1, len(v))
		for i, value := range v {
			result[i], _ = AsNestedStruct1(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []NestedStruct1", v, v)
	}
}
func AsNestedStruct2(v any) (NestedStruct2, error) {
	switch v := v.(type) {
	case NestedStruct2:
		return v, nil
	default:
		return NestedStruct2{}, fmt.Errorf("unable to cast %#v of type %T to NestedStruct2", v, v)
	}
}

func AsNestedStruct2Array(v any) ([]NestedStruct2, error) {
	switch v := v.(type) {
	case []NestedStruct2:
		return v, nil
	case []interface{}:
		result := make([]NestedStruct2, len(v))
		for i, value := range v {
			result[i], _ = AsNestedStruct2(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []NestedStruct2", v, v)
	}
}
func AsNestedStruct3(v any) (NestedStruct3, error) {
	switch v := v.(type) {
	case NestedStruct3:
		return v, nil
	default:
		return NestedStruct3{}, fmt.Errorf("unable to cast %#v of type %T to NestedStruct3", v, v)
	}
}

func AsNestedStruct3Array(v any) ([]NestedStruct3, error) {
	switch v := v.(type) {
	case []NestedStruct3:
		return v, nil
	case []interface{}:
		result := make([]NestedStruct3, len(v))
		for i, value := range v {
			result[i], _ = AsNestedStruct3(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []NestedStruct3", v, v)
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
