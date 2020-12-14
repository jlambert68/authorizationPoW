
# cat -e -t -v Makefile

#.DEFAULT_GOAL := buildAndRun

compileProto:
	@echo "Compile proto file..."
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I ../grpc_api/secretMessageGenerator_grpc_api --go_out=plugins=grpc:../grpc_api/secretMessageGenerator_grpc_api ../grpc_api/secretMessageGenerator_grpc_api/secretMessageGenerator_grpc_api.proto


