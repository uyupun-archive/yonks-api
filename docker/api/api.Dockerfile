FROM golang:1.17.3 AS builder
RUN mkdir /go/src/yonks
WORKDIR /go/src/yonks
COPY ./ .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o bin/yonks cmd/main.go cmd/config.go cmd/router.go
RUN go build -o bin/migrator migrations/main.go migrations/targets.go
RUN go build -o bin/seeder seeds/main.go seeds/targets.go

FROM alpine:3.15.0
WORKDIR /app/
COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo
COPY --from=builder /go/src/yonks/bin/yonks .
COPY --from=builder /go/src/yonks/bin/migrator .
COPY --from=builder /go/src/yonks/bin/seeder .
COPY .env ./.env
COPY ./seeds/init ./seeds/init
COPY ./seeds/mock ./seeds/mock
ENTRYPOINT ["./yonks"]
