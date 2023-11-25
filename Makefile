DOCKER_IMAGE_NAME = user-service
DOCKER_IMAGE_TAG = 1.0.0

.PHONY: protoc
protoc: ## Generate protobuf files
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.18.0
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	@echo "Generating protobuf files..."
	@protoc -I ./proto --go_out=./ \
	--go-grpc_out=require_unimplemented_servers=false:./ \
	--grpc-gateway_out . --grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt generate_unbound_methods=true \
	./proto/users/*.proto ./proto/google/*.proto
	@echo "Done"
	
.PHONY: # build docker image
docker:
	@echo "build docker image and deploy user service..."
	@docker build -t $(DOCKER_IMAGE_NAME) .
	@docker tag $(DOCKER_IMAGE_NAME) $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

.PHONY: run
run: ## Run the service
	@echo "Run the service..."
	@go run main.go