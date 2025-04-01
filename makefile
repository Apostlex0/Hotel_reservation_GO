build :
	@go build main.go

run: build 
	@go run main.go

test :
	go test ./...