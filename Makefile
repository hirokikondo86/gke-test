protoc-client:
	sh ./script/protoc.sh
	sh ./script/disable.sh
protoc-server:
	protoc -I protoc/ protoc/enums.proto --go_out=plugins=grpc:.
	protoc -I protoc/ protoc/hello_messages.proto --go_out=plugins=grpc:.
	protoc -I protoc/ protoc/hello_services.proto --go_out=plugins=grpc:.
protoc-doc:
	protoc \
	--doc_out=./protoc/doc \
	--doc_opt=html,index.html:./protoc protoc/*.proto

# realize start
start-server:
	docker exec -it ecs-test-server bash -c "realize start"
# go build
build-server:
	docker exec -it ecs-test-server bash -c "CGO_ENABLED=0 &&\
	GOOS=linux &&\
	GOARCH=amd64 &&\
	cd src &&\
	go build -o ../bin/main"

# dep init
dep-init:
	docker exec -it ecs-test-server dep init
# dep ensure
dep-ensure:
	docker exec -it ecs-test-server dep ensure