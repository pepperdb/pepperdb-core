PB = $(wildcard *.proto)
GO = $(PB:.proto=.pb.go)

all: $(GO)

%.pb.go: %.proto
	protoc -I/usr/local/include -I. -I../../../../../../src --gogo_out=. $<

clean:
	rm *.pb.go
