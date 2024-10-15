LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

#./... рекурсивно пройтись по всем папкам, --config для кастомизации
lint:
	golangci-lint run ./... --config .golangci.pipeline.yaml
#$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml


# protoc-gen-go будет скачан в локальную папку bin 
#LOCAL_BIN указывает куда будет скачан
install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

# для генерации нескольких апишек последовательно за одну команду
generate:
	make generate-chat-api


generate-chat-api:
	mkdir -p pkg/chat_v1
	protoc --proto_path api/chat_v1 \
	--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/chat_v1/chat.proto

#--proto_path api/chat_v1 передаём путь до прото файла
#--go_out=pkg/chat_v1 хотим сгенерировать гошные структуры данных
#сам protoc не генерит гошный файл, он роутит декларацию протофбафа к бинарнику protoc-gen-go 
# --go_opt=paths=source_relative опция для того чтобы соориентировать куда нужно сохронять сгенерированный код
# --plugin=protoc-gen-go=bin/protoc-gen-go указывает на то что protoc-gen-go нужно искать локально в bin
# бинарник protoc-gen-go-grpc нужен для генерации сервера и клиента из прото файла(декларация)
# бинарник protoc-gen-go-grpc также будет лежать в bin
# api/chat_v1/chat.proto прямой путь до протофайла

# protoc-gen-go будет скачан в локальную папку bin 


build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/grpc_server/main.go

copy-to-server:
	scp -i /mnt/c/Users/titva/.ssh/id_ed25519_for_selectel service_linux root@176.114.77.183:

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/chat-server/test-server:v0.0.1 .
	docker login -u token -p CRgAAAAA7KqVel9s8FyxfwCa5JKAZ3Z9N8MeJbe1 cr.selcloud.ru/chat-server
	docker push cr.selcloud.ru/chat-server/test-server:v0.0.1

