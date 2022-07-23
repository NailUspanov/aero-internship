create:
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/
	protoc -I ./proto --grpc-gateway_out ./gen/ \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
        proto/news.proto


clean:
	rm gen/*.go

build:
	go mod tidy && go build -o app.out ./cmd/app

run: build
	./app.out