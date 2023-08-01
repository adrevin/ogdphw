BIN_CALENDAR := "./bin/calendar"
BIN_SCHEDULER := "./bin/scheduler"
BIN_SENDER := "./bin/sender"
DOCKER_CALENDAR_IMG := "calendar:develop"
DOCKER_SCHEDULER_IMG := "scheduler:develop"
DOCKER_SENDER_IMG := "sender:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN_CALENDAR) -ldflags "$(LDFLAGS)" ./cmd/calendar
	go build -v -o $(BIN_SCHEDULER) -ldflags "$(LDFLAGS)" ./cmd/scheduler
	go build -v -o $(BIN_SENDER) -ldflags "$(LDFLAGS)" ./cmd/sender

run: build
	$(BIN_CALENDAR) -config ./configs/config.yml

build-calendar-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_CALENDAR_IMG) \
		-f build/calendar/Dockerfile .

build-scheduler-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_SCHEDULER_IMG) \
		-f build/scheduler/Dockerfile .

build-sender-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_SENDER_IMG) \
		-f build/sender/Dockerfile .

build-img: build-calendar-img build-scheduler-img build-sender-img

run-img: build-img
	docker run $(DOCKER_CALENDAR_IMG)

version: build
	$(BIN_CALENDAR) version

test:
	go test -race ./internal/... ./pkg/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.52.2

lint: install-lint-deps
	golangci-lint run ./...

generate:
	go generate ./internal/server/grpc/.

up:
	docker compose -f ./deployments/docker-compose.yml up -d --remove-orphans

down:
	docker compose -f ./deployments/docker-compose.yml down

integration-tests:
	go test ./integration/. -v -count=1 -tags=integration_tests

.PHONY: build run build-img run-img version test lint
