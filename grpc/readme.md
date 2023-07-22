## 1、protobuf

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-gr
pc_opt=paths=source_relative pb/greeter.proto