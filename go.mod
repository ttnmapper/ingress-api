module ttnmapper-ingress-api

go 1.18

// TTN's fork of throttled/throttled/v2.
replace github.com/throttled/throttled/v2 => github.com/TheThingsIndustries/throttled/v2 v2.7.1-noredis

require (
	github.com/766b/chi-prometheus v0.0.0-20180509160047-46ac2b31aa30
	github.com/brocaar/chirpstack-api/go/v3 v3.11.1
	github.com/chirpstack/chirpstack/api/go/v4 v4.6.0
	github.com/go-chi/chi v3.3.3+incompatible
	github.com/go-chi/render v1.0.1
	github.com/golang/protobuf v1.5.3
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.14.0
	github.com/streadway/amqp v1.0.0
	github.com/tkanos/gonfig v0.0.0-20210106201359-53e13348de2f
	github.com/ulule/deepcopier v0.0.0-20200430083143-45decc6639b6
	go.thethings.network/lorawan-stack/v3 v3.24.2-0.20230303154842-4ac6c7ddca34
	golang.org/x/net v0.17.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/TheThingsIndustries/protoc-gen-go-flags v1.0.6 // indirect
	github.com/TheThingsIndustries/protoc-gen-go-json v1.4.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.2 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gotnospirit/makeplural v0.0.0-20180622080156-a5f48d94d976 // indirect
	github.com/gotnospirit/messageformat v0.0.0-20221001023931-dfe49f1eb092 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/grpc v1.59.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
