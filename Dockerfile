FROM golang:1.22-alpine AS builder

RUN apk add curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/mole-squad/soq-api
COPY . .

RUN task build

FROM alpine:latest

COPY --from=builder /go/src/mole-squad/soq-api/bin/soq /bin/soq

RUN mkdir -p /app/.profile.d/ && \
  echo '[ -z "$SSH_CLIENT" ] && source <(curl --fail --retry 3 -sSL "$HEROKU_EXEC_URL")' > /app/.profile.d/heroku-exec.sh && \
  rm /bin/sh && ln -s /bin/bash /bin/sh

CMD ["/bin/soq", "api"]
