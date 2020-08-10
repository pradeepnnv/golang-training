Reference: https://github.com/confluentinc/confluent-kafka-go#installing-librdkafka

Since bundled librdkafka is not supported, we need to build the application and then execute it.

```
go build -tags dynamic
```

