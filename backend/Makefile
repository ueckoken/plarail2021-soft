all:statesync
	
statesync:statesync_internal statesync_external

statesync_internal:
	protoc \
	--go_out=internal \
	--go_opt=Mproto/statestnc.proto=. \
	--go-grpc_out=internal \
	--go-grpc_opt=Mproto/statesync.proto=. \
	-I./ \
	proto/statesync.proto
statesync_external:
	protoc \
	--go_out=external \
	--go_opt=Mproto/statestnc.proto=. \
	--go-grpc_out=external \
	--go-grpc_opt=Mproto/statesync.proto=. \
	-I./ \
	proto/statesync.proto
