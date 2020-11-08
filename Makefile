protos:
	mkdir -p ./src/api/servicepb
	protoc protos/service.proto -I protos/ -I protos/vendors/googleapis/ -I protos/vendors/grpc-gateway/ \
	  --go_out ./src/api/servicepb/ --go_opt paths=source_relative \
	  --go-grpc_out ./src/api/servicepb/ --go-grpc_opt paths=source_relative \
	  --grpc-gateway_out ./src/api/servicepb/ --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative \
	  --openapiv2_out ./ --openapiv2_opt logtostderr=true

http_client:
	swagger generate client -f ./service.swagger.json -A service --target src/http_client

.PHONY: protos
