package main

{{- $system := .System }}
import (
    "context"
	"flag"
	"log"
	"net/http"
    "github.com/apigear-io/objectlink-core-go/olink/ws"
    "github.com/apigear-io/objectlink-core-go/olink/remote"
{{- range .System.Modules}}
{{- $module := . }}
    {{snake .Name}}_olink "{{$system.Name}}/{{snake .Name}}/olink"
    {{snake .Name}}_impl "{{$system.Name}}/{{snake .Name}}/olink/testbed"
{{- end }}    
)


var addr = flag.String("addr", ":8080", "http service address")

var registry = remote.NewRegistry()
var node = remote.NewNode(registry)
var ctx = context.Background()
var hub = ws.NewHub(ctx, registry)

func init() {
{{- range .System.Modules }}
{{- $module := . }}    
{{- range .Interfaces }}
    {   // register {{ snake $module.Name}} module
        source := {{snake $module.Name}}_olink.New{{Camel .Name}}Source()
        registry.AddObjectSource(source)
        impl := {{snake $module.Name}}_impl.New{{Camel .Name}}(source)
        source.SetImplementation(impl)
        registry.LinkRemoteNode(source.ObjectId(), node)
    }
{{- end }}
{{- end}}
}

func main() {
	flag.Parse()
    http.HandleFunc("/ws", hub.ServeHTTP)
    log.Printf("listen on %s/ws", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}