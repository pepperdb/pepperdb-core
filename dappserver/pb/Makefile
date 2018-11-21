PB = $(wildcard *.proto)
GO = $(PB:.proto=.pb.go)

all: $(GO)

config.pb.go: config.proto
	protoc --gogo_out=. $<
	sed 's/,omitempty//g' config.pb.go > tmp
	mv tmp  config.pb.go

grpc.pb.go: grpc.proto
	protoc -I/usr/local/include -I. -I../../../../../../src -I../../vendor/github.com/nebulasio/grpc-gateway/third_party/googleapis --gogo_out=plugins=grpc:. $<

clean:
	rm *.pb.go