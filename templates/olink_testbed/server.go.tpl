package main

{{- $system := .System }}
import (
	"flag"
	"log"
	"net/http"
    "olink/pkg/ws"
    "olink/pkg/remote"
{{- range .System.Modules}}
{{- $module := . }}
    {{snake .Name}}_olink "{{$system.Name}}/{{snake .Name}}/olink"
    {{snake .Name}}_impl "{{$system.Name}}/{{snake .Name}}/olink/testbed"
{{- end }}    
)


var addr = flag.String("addr", ":8080", "http service address")

var registry = remote.NewRegistry()
var node = remote.NewNode(registry)
var hub = ws.NewHub(registry)

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