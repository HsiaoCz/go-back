## 1、protobuf

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-gr
pc_opt=paths=source_relative pb/greeter.proto

protoc -I=v1 \ master ✱
--go_out=v1 --go_opt=paths=source_relative \
 --go-grpc_out=v1 --go-grpc_opt=paths=source_relative \
 --grpc-gateway_out=v1 --grpc-gateway_opt=paths=source_relativ
e \
 v1/greeter.proto
