FROM golang:1.22-alpine AS builder

RUN apk add curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/burkel24/taskapp
COPY . .

RUN task taskapp_build
RUN ls  ./bin

FROM alpine:latest

COPY --from=builder /go/src/burkel24/taskapp/bin/taskapp /bin/taskapp
CMD ["/bin/taskapp", "api"]
