PB = $(wildcard *.proto)
GO = $(PB:.proto=.pb.go)

all: $(GO)

%.pb.go: %.proto
	protoc -I/usr/local/include -I. -I../../../../../../../src -I../../../vendor/github.com/nebulasio/grpc-gateway/third_party/googleapis --gogo_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. $<
	sed 's/grpc-ecosystem/nebulasio/g' rpc.pb.gw.go > tmp
	mv tmp rpc.pb.gw.go
clean:
	rm *.go
