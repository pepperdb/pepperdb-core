PB = $(wildcard *.proto)
GO = $(PB:.proto=.pb.go)

all: $(GO)

%.pb.go: %.proto
	protoc --gogo_out=. $<
	sed 's/,omitempty//g' config.pb.go > tmp
	mv tmp  config.pb.go
clean:
	rm *.pb.go