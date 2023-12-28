FROM golang:latest as builder

WORKDIR /go-modules

COPY . ./

# Building using -mod=vendor, which will utilize the v
#RUN go get
#RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -v -mod=vendor -o ingress-api -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn"

#FROM alpine:3.8
FROM scratch

WORKDIR /root/

COPY --from=builder /go-modules/ingress-api .
COPY conf.json .

CMD ["./ingress-api"]