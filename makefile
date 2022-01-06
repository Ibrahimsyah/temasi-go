generate-proto:
	protoc --go_out=plugins=grpc:. ./shared/protos/*.proto 