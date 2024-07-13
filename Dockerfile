FROM golang:1.22-alpine AS builder

RUN apk add curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/mole-squad/soq-api
COPY . .

RUN task build

FROM alpine:latest

COPY --from=builder /go/src/mole-squad/soq-api/bin/soq /bin/soq

CMD ["/bin/soq", "api"]
