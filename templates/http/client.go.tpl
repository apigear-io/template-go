package http


import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type Client struct {
    client *http.Client
}

func NewClient(c *http.Client) *Client {
    return &Client{
        client: c,
    }
}

func (c *Client) Post(url string, data interface{}) (interface{}, error) {
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

{{- range .Module.Interfaces }}
{{- $iface := .Name }}

{{- range .Operations }}

{{- $name := (print $iface (Camel .Name))}}

func (c *Client) {{ $name}}(req *{{ $name }}Request, reply *{{ $name }}Reply) error {
    url := "{{ $.Module.Name }}/{{ $iface }}/{{ .Name }}"
    resp, err := c.Post(url, req)
    if err != nil {
        return err
    }
    *reply = *resp.(*{{ $name }}Reply)
    return nil
}

{{ end }}
{{ end }}
