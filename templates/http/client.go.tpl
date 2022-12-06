package {{.Module.Name}}


import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

{{- range .Module.Interfaces }}
{{- $class := printf "%sClient" (Camel .Name) }}

type {{ $class}} struct {
    baseUrl string
    client *http.Client
}

func New{{$class}}(baseUrl string, c *http.Client) *{{$class}} {
    return &{{$class}}{
        baseUrl: baseUrl,
        client: c,
    }
}

func (c *{{$class}}) post(url string, data interface{}) (interface{}, error) {
    b, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }
    resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(b))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("%s", resp.Status)
    }
    var result interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }
    return result, nil
}
{{- $iface := .Name }}

{{- range .Operations }}

{{- $name := (print $iface (Camel .Name))}}

func (c *{{$class}}) {{Camel .Name}}(req *{{ $name }}Request) (*{{ $name }}Reply, error) {
    resp, err := c.post("{{ snake $.Module.Name }}/{{ snake $iface }}/{{ snake .Name }}", req)
    if err != nil {
        return nil, err
    }
    reply := resp.(*{{ $name }}Reply)
    return reply, nil
}
{{- end }}
{{- end }}
