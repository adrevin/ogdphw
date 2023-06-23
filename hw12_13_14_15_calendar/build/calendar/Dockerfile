FROM golang:1.20 as build

ENV BIN_FILE /opt/calendar/calendar-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build -ldflags "$LDFLAGS" -o ${BIN_FILE} cmd/calendar/*

FROM alpine:3.9

LABEL ORGANIZATION="OTUS Online Education"
LABEL SERVICE="calendar"
LABEL MAINTAINERS="student@otus.ru"

ENV BIN_FILE "/opt/calendar/calendar-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/calendar/config.yml
COPY --from=build /go/src/configs/config.yml ${CONFIG_FILE}

ENV MIGRATIONS_FOLDER "/opt/calendar/migrations/"
COPY --from=build /go/src/migrations/*  ${MIGRATIONS_FOLDER}

CMD ${BIN_FILE} -config ${CONFIG_FILE} migrate-db; ${BIN_FILE} -config ${CONFIG_FILE}
