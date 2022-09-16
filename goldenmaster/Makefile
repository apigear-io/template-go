.PHONY: format testbed testserver


format:
	gofmt -w .
	goimports -w .
	go mod tidy

test: format
	go test -v goldenmaster/tb_adv/olink/testbed
	go test -v goldenmaster/tb_conflict/olink/testbed
	go test -v goldenmaster/tb_data/olink/testbed
	go test -v goldenmaster/tb_enum/olink/testbed
	go test -v goldenmaster/tb_same/olink/testbed
	go test -v goldenmaster/tb_again/olink/testbed
	go test -v goldenmaster/tb_simple/olink/testbed


server: format
	go run cmd/olink_testbed_server/main.go

debug-server: format
	dlv debug cmd/olink_testbed_server/main.go	