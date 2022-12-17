# PROTO_DIR=protoopt
# PROTO_DIR=protoout
# PACKAGE=shiva
# GOOUT=.
# BIN=binfiles

#models and services seperated
type:
	# best models(book user) , services (book users)
	protoc --proto_path=type-driven --go_out=. --go-grpc_out=. --go_opt=module=shiva --go-grpc_opt=module=shiva type-driven/*.proto

#seperated by domain
domain:
	# works fine will have book (models and services), users (models and services)
	protoc --proto_path=domain-driven --go_out=. --go-grpc_out=. --go_opt=module=shiva --go-grpc_opt=module=shiva domain-driven/*.proto

#does not work
out:
	# does not work (will get importing error in models amd services)
	protoc --proto_path=protoout --go_out=codeout/models --go-grpc_out=codeout/services --go_opt=module=shiva --go-grpc_opt=module=shiva protoout/*.proto

# encode:
# 	protoc --encode=One ${PROTO_DIR}/one.proto > ${BIN}/one.bin 
# decode:
# 	cat ${BIN}/one.bin | protoc --decode=One ${PROTO_DIR}/one.proto > ${BIN}/one.txt




