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

