
protoc -I=$PWD --go_out=$PWD/pb --go-grpc_out=$PWD/pb $PWD/protos/*.proto
