FROM golang:1.20 as build

ENV BIN_FILE /opt/sender/sender-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build -ldflags "$LDFLAGS" -o ${BIN_FILE} cmd/sender/*

FROM alpine:3.9

LABEL ORGANIZATION="OTUS Online Education"
LABEL SERVICE="sender"
LABEL MAINTAINERS="student@otus.ru"

ENV BIN_FILE "/opt/sender/sender-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/sender/config.yml
COPY --from=build /go/src/configs/config.yml ${CONFIG_FILE}


CMD ${BIN_FILE} -config ${CONFIG_FILE}
