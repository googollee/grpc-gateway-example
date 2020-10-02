protos:
	protoc -I ./protos protos/service.proto --go_out=plugins=grpc:./backend/servicepb/
	protoc -I ./protos protos/service.proto --go_out=plugins=grpc:./client/servicepb/
	protoc -I ./protos protos/service.proto --grpc-gateway_out ./backend/servicepb --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative
	protoc -I ./protos protos/service.proto --swagger_out ./api/ --swagger_opt logtostderr=true

.PHONY: protos
