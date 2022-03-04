FROM golang:1.17

WORKDIR /app
COPY . .
RUN go generate
RUN go mod download
RUN go build

CMD [ "./machine-producer" ]
