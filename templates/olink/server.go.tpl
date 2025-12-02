package main

{{- $system := .System }}
{{ $import := .System.Meta.GetString "go.module" }}
import (
    "context"
	"flag"
	"net/http"
    "os"

    "github.com/apigear-io/objectlink-core-go/olink/ws"
    "github.com/apigear-io/objectlink-core-go/olink/remote"
    "github.com/rs/zerolog/log"
    "github.com/rs/zerolog"
{{- range .System.Modules}}
{{- $module := . }}
    "{{$import}}/{{snake .Name}}"
{{- end }}
)


var addr = flag.String("addr", "127.0.0.1:5555", "http service address")
var registry = remote.NewRegistry()
var ctx = context.Background()
var hub = ws.NewHub(ctx, registry)

func init() {
{{- range .System.Modules }}
{{- $module := snake .Name }}

{{- range .Interfaces }}
{{- $type := Camel .Name }}
    {   // register {{ $module }} module
        source := {{ $module }}.New{{$type}}Source()
        registry.AddObjectSource(source)
        impl := {{ $module }}.NewDefault{{$type}}(source)
        source.SetImplementation(impl)
    }
{{- end }}
{{- end}}
}

func main() {
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	flag.Parse()
    http.HandleFunc("/ws", hub.ServeHTTP)
    log.Info().Msgf("ObjectLink server started at ws://%s/", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Error().Msgf("ListenAndServe: %s", err)
        return
	}
}
