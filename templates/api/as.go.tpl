package api

import (
	"encoding/json"
	"fmt"
)


func AsInt(v any) (int32, error) {
	switch v := v.(type) {
	case int32:
		return v, nil
	case json.Number:
		i, err := v.Int64()
        if err != nil {
            return 0, err
        }
        return int32(i), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", v, v)
	}
}

func AsInt64(v any) (int64, error) {
	switch v := v.(type) {
	case int64:
		return v, nil
	case json.Number:
		i, err := v.Int64()
		if err != nil {
			return 0, err
		}
		return i, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", v, v)
	}
}



func AsIntArray(v any) ([]int32, error) {
	switch v := v.(type) {
	case []int32:
		return v, nil
	case []interface{}:
		result := make([]int32, len(v))
		for i, value := range v {
			result[i], _ = AsInt(value)
		}
		return result, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int32", v, v)
	}
}

func AsInt32(v any) (int32, error) {
	return AsInt(v)
}

func AsInt32Array(v any) ([]int32, error) {
	return AsIntArray(v)
}

func AsFloat(v any) (float32, error) {
	switch v := v.(type) {
	case float32:
		return v, nil
	case json.Number:
		f, err := v.Float64()
		if err != nil {
			return 0.0, err
		}
		return float32(f), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", v, v)
	}
}



func AsFloatArray(v any) ([]float32, error) {
	switch v := v.(type) {
	case []float32:
		return v, nil
	case []interface{}:
		result := make([]float32, len(v))
		for i, value := range v {
			result[i], _ = AsFloat(value)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []float32", v, v)
	}
}

func AsFloat32(v any) (float32, error) {
	return AsFloat(v)
}


func AsFloat32Array(v any) ([]float32, error) {
	return AsFloatArray(v)
}

func AsFloat64(v any) (float64, error) {
	switch v := v.(type) {
	case float64:
		return v, nil
	case json.Number:
		f, err := v.Float64()
		if err != nil {
			return 0.0, err
		}
		return f, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", v, v)
	}
}


func AsFloat64Array(v any) ([]float64, error) {
	switch v := v.(type) {
	case []float64:
		return v, nil
	case []interface{}:
		result := make([]float64, len(v))
		for i, value := range v {
			result[i], _ = AsFloat64(value)
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
	case json.Number:
		i, err := v.Int64()
		if err != nil {
			return false, err
		}
		return i != 0, nil
	case string:
		return v == "true", nil
	case int:
		return v != 0, nil
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

{{- range .Module.Interfaces }}
func As{{Camel .Name}}(v any) ({{.Name}}, error) {
    switch v := v.(type) {
    case {{.Name}}:
        return v, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to {{.Name}}", v, v)
    }
}

func As{{Camel .Name}}Array(v any) ([]{{.Name}}, error) {
    switch v := v.(type) {
    case []{{.Name}}:
        return v, nil
    case []interface{}:
        result := make([]{{.Name}}, len(v))
        for i, value := range v {
            result[i], _ = As{{.Name}}(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []{{.Name}}", v, v)
    }
}
{{- end }}

{{- range .Module.Enums }}
{{- end }}

{{- range .Module.Structs }}
func As{{Camel .Name}}(v any) ({{.Name}}, error) {
    switch v := v.(type) {
	case map[string]any:
		result := {{.Name}}{}
		data, err := json.Marshal(v)
		if err != nil {
			return {{.Name}}{}, err
		}
		err = json.Unmarshal(data, &result)
		if err != nil {
			return {{.Name}}{}, err
		}
		return result, nil
    case {{.Name}}:
        return v, nil
    default:
        return {{.Name}}{}, fmt.Errorf("unable to cast %#v of type %T to {{.Name}}", v, v)
    }
}

func As{{Camel .Name}}Array(v any) ([]{{.Name}}, error) {
    switch v := v.(type) {
    case []{{.Name}}:
        return v, nil
    case []interface{}:
        result := make([]{{.Name}}, len(v))
        for i, value := range v {
            result[i], _ = As{{.Name}}(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []{{.Name}}", v, v)
    }
}
{{- end }}

{{- range .Module.Enums }}
func As{{Camel .Name}}(v any) ({{.Name}}, error) {
    switch v := v.(type) {
    case {{.Name}}:
        return v, nil
    default:
        return {{.Name}}_UNKNOWN, fmt.Errorf("unable to cast %#v of type %T to {{.Name}}", v, v)
    }
}

func As{{Camel .Name}}Array(v any) ([]{{.Name}}, error) {
    switch v := v.(type) {
    case []{{.Name}}:
        return v, nil
    case []interface{}:
        result := make([]{{.Name}}, len(v))
        for i, value := range v {
            result[i], _ = As{{.Name}}(value)
        }
        return result, nil
    default:
        return nil, fmt.Errorf("unable to cast %#v of type %T to []{{.Name}}", v, v)
    }
}
{{- end }}
