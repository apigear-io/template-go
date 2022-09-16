package main
import (
    "flag"
	"log"
	"olink/pkg/client"
	"olink/pkg/ws"    
)

var addr = flag.String("addr", "ws://127.0.0.1:8080/ws", "ws service addr")


var registry = client.NewRegistry()
var node = client.NewNode(registry)


func registerSinks() {
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