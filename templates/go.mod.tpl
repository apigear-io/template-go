module {{.System.Name}}

go 1.18

require olink v0.0.0

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
)

replace olink => ../../objectlink-core-go/