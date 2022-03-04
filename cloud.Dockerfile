FROM golang:1.17

RUN apt update && \
    apt install unzip

# Download proto zip
ENV PROTOC_ZIP=protoc-3.14.0-linux-x86_64.zip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/${PROTOC_ZIP}
RUN unzip -o ${PROTOC_ZIP} -d ./proto
RUN chmod 755 -R ./proto/bin
ENV BASE=/usr
# Copy into path
RUN cp ./proto/bin/protoc ${BASE}/bin/
RUN cp -R ./proto/include/* ${BASE}/include/
# install go protocol buffers
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

WORKDIR /app/factory-shared
COPY factory-shared/* ./
RUN go generate

WORKDIR /app/factory-cloud
COPY factory-cloud/* ./

RUN go mod download
RUN go build

CMD [ "./factory-cloud" ]
