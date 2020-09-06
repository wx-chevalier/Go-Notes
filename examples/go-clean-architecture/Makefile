.Phony: begin
begin:
	@~/.air -d -c .air.conf

.phony: build
build:
	CGOENABLED=0 go build -o bin/main api/main.go
