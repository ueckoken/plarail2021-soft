up:
	docker-compose build
	docker-compose up
pb: statesync speedControl
	
statesync: statesync_internal statesync_external

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

speedControl: speedControl_builtin speedControlSpeed

#speedControl_internal:
#	protoc \
#	--go_out=internal/spec \
#	--go_opt=Mproto/speedControl.proto=. \
#	--go-grpc_out=internal \
#	--go-grpc_opt=Mproto/speedControl.proto=. \
#	-I./ \
#	proto/speedControl.proto
#
#speedControl_external:
#	protoc \
#	--go_out=external/spec \
#	--go_opt=Mproto/speedControl.proto=. \
#	--go-grpc_out=external \
#	--go-grpc_opt=Mproto/speedControl.proto=. \
#	-I./ \
#	proto/speedControl.proto

speedControl_builtin:
	protoc \
	--go_out=builtin/speed/spec \
	--go_opt=Mproto/speedControl.proto=. \
	--go-grpc_out=builtin/speed \
	--go-grpc_opt=Mproto/speedControl.proto=. \
	-I./ \
	proto/speedControl.proto

speedControlSpeed:
	protoc \
	--go_out=speed/spec \
	--go_opt=Mproto/speedControl.proto=. \
	--go-grpc_out=speed \
	--go-grpc_opt=Mproto/speedControl.proto=. \
	-I./ \
	proto/speedControl.proto
