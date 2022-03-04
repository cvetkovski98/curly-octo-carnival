module github.com/cvetkovski98/factory-cloud

go 1.17

replace github.com/cvetkovski98/factory-shared => ../factory-shared

require github.com/nats-io/nats.go v1.13.1-0.20220121202836-972a071d373d

require (
	github.com/cvetkovski98/factory-shared v0.0.0
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/nats-io/nats-server/v2 v2.7.2 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	google.golang.org/protobuf v1.27.1
)
