package main
{{- $system := .System }}
import (
    "flag"
	"log"
	"olink/pkg/client"
	"olink/pkg/ws"
{{- range .System.Modules}}
{{- $module := . }}
    {{snake .Name}}_olink "{{$system.Name}}/{{snake .Name}}/olink"
{{- end }}    
)

var addr = flag.String("addr", "ws://127.0.0.1:8080/ws", "ws service addr")


var registry = client.NewRegistry()
var node = client.NewNode(registry)


func registerSinks() {
{{- range .System.Modules }}	
{{- $module := . }}
{{- range .Interfaces }}
	{ // register {{ snake $module.Name}} module
		sink := {{snake $module.Name}}_olink.New{{.Name}}Sink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
{{- end }}
{{- end }}
}

func main() {
    flag.Parse()
    conn, err := ws.Dial(*addr)
    	if err != nil {
		log.Fatalf("dial error: %s\n", err)
		return
	}
    defer conn.Close()
	node.SetOutput(conn)
	conn.SetOutput(node)
	registry.AttachClientNode(node)
	registerSinks()


}