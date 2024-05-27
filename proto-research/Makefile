gen:
	protoc --proto_path=./proto \
        --go_out=./gen \
        --go-grpc_out=./gen \
        --grpc-gateway_out=./gen \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt generate_unbound_methods=true

.PHONY: gen