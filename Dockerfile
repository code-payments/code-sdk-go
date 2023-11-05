FROM golang:1.21

RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1

WORKDIR /sdk

CMD ["bash"]
