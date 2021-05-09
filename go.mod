module ttnmapper-ingress-api

go 1.16

replace gopkg.in/DATA-DOG/go-sqlmock.v1 => gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0

require (
	github.com/766b/chi-prometheus v0.0.0-20180509160047-46ac2b31aa30
	github.com/brocaar/chirpstack-api/go/v3 v3.9.7
	github.com/go-chi/chi v3.3.3+incompatible
	github.com/go-chi/render v1.0.1
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1
	github.com/streadway/amqp v0.0.0-20181205114330-a314942b2fd9
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	github.com/ulule/deepcopier v0.0.0-20200117111125-792cfb847af8
	go.thethings.network/lorawan-stack/v3 v3.12.3
	golang.org/x/net v0.0.0-20201021035429-f5854403a974
)
