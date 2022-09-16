package main
import (
	"flag"
	"log"
	"net/http"
    "olink/pkg/ws"
    "olink/pkg/remote"    
)


var addr = flag.String("addr", ":8080", "http service address")

var registry = remote.NewRegistry()
var node = remote.NewNode(registry)
var hub = ws.NewHub(registry)

func init() {
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