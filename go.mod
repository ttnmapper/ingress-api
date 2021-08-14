module ttnmapper-ingress-api

go 1.16

replace gopkg.in/DATA-DOG/go-sqlmock.v1 => gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0

require (
	github.com/766b/chi-prometheus v0.0.0-20180509160047-46ac2b31aa30
	github.com/brocaar/chirpstack-api/go/v3 v3.11.1
	github.com/envoyproxy/protoc-gen-validate v0.6.1 // indirect
	github.com/go-chi/chi v3.3.3+incompatible
	github.com/go-chi/render v1.0.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/common v0.30.0 // indirect
	github.com/prometheus/procfs v0.7.2 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/tkanos/gonfig v0.0.0-20210106201359-53e13348de2f
	github.com/ulule/deepcopier v0.0.0-20200430083143-45decc6639b6
	github.com/vmihailenco/msgpack/v5 v5.3.4 // indirect
	go.thethings.network/lorawan-stack/v3 v3.14.1
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210813162853-db860fec028c // indirect
	google.golang.org/grpc v1.40.0 // indirect
)
