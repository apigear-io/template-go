package main
{{- $system := .System }}
{{ $import := .System.Meta.GetString "go.module" }}
import (
    "flag"
	"context"
	"os"
	"github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/ws"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
{{- range .System.Modules}}
    "{{$import}}/{{snake .Name}}"
{{- end }}
)


var registry = client.NewRegistry()
var node = client.NewNode(registry)


func registerSinks() {
{{- range .System.Modules }}
{{- $module := . }}
{{- range .Interfaces }}
	{ // register {{ snake $module.Name}} module scope
		sink := {{snake $module.Name}}.New{{.Name}}Sink(node)
		registry.AddObjectSink(sink)
		node.LinkRemoteNode(sink.ObjectId())
	}
{{- end }}
{{- end }}
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	flag.Usage = func() {
		log.Info().Msgf("Usage: %s [options]", os.Args[0])
		flag.PrintDefaults()
	}
	var addr = flag.String("addr", "ws://127.0.0.1:5555/ws", "ws service addr")
    flag.Parse()
	ctx := context.Background()
    conn, err := ws.Dial(ctx, *addr)
    	if err != nil {
		log.Error().Err(err).Msg("ws dial error")
		return
	}
    defer conn.Close()
	node.SetOutput(conn)
	conn.SetOutput(node)
	registry.AttachClientNode(node)
	registerSinks()

	select {} // block forever

}
