package {{.Module.Name}}

import (
    "net/http"

    "github.com/labstack/echo/v4"
)




{{- range .Module.Interfaces }}
{{- $iface := .Name }}
{{- $class := printf "%sHandler" .Name }}

type {{$class}} struct {
    s I{{.Name}}Service    
}

func New{{$class}}(s I{{.Name}}Service) *{{$class}} {
    return &{{$class}}{s}
}

{{- range .Operations }}
{{- $name := printf "%s%s" (Camel $iface) (Camel .Name) }}

func (h *{{$class}}) Handle{{Camel .Name}}(c echo.Context) error {
    var req {{$name}}Request
    err := c.Bind(&req)
    if err != nil {
        return err
    }
    resp, err := h.s.{{Camel .Name}}(&req)
    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, resp)
}
{{- end }}
{{- end }}
