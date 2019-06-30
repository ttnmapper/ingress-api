# TTN Mapper Ingress API
REST API where new raw measurement data will be received.


# Configuration

Configuration can be passed as environment variables, or by setting the values in config.json. For the latter copy config.json.template and edit the values.

| JSON key | Type | ENV var key | Description |
| --- | --- | --- | --- |
| AmqpHost     | string | AMQP_HOST | The hostname for the server where RabbitMQ is running.
| AmqpPort     | string | AMQP_PORT | The port on which RabbitMQ is listening.
| AmqpUser     | string | AMQP_USER | RabbitMQ username.
| AmqpPassword | string | AMQP_PASSWORD | RabbitMQ password.


# Endpoints

## /ttn

### GET /ttn/v2
Returns with `{"message": "GET test success"}`
### POST /ttn/v2
Endpoint to which TTNv2 and TTIv2 stacks should POST data.

### GET /ttn/v3
Returns with `{"message": "GET test success"}`

### POST /ttn/v3
Not implemented yet.

## /android
### POST /android/v2/
Endpoint for data in the format used by https://github.com/ttnmapper/ttnmapper_android_v2

### POST /android/v3/
Endpoint for data in the format used by https://github.com/ttnmapper/ttnmapper-android-v3

## /ios
### POST /ios/v2
Forwards to /android/v2 as it is the same format.