.PHONY: format testbed testserver


format:
	gofmt -w .
	goimports -w .
	go mod tidy

test: format
{{- $name := .System.Name}}
{{- range .System.Modules }}
	go test -v {{$name}}/{{ snake .Name }}/olink/testbed
{{- end }}


server: format
	go run cmd/olink_testbed_server/main.go

debug-server: format
	dlv debug cmd/olink_testbed_server/main.go	