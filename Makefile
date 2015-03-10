prod:
	go build main.go

debug:
	go build -tags=debug main.go
