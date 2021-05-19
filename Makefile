VERSION     = $(shell /usr/bin/git describe --always)
BUILD_DATE  = $(shell date +%s)

GO_LDFLAGS  = -X main.version=$(if $(VERSION),$(VERSION),unknown)
GO_LDFLAGS += -X main.buildDate=$(if $(BUILD_DATE),$(BUILD_DATE),unknown)

build:
	go build -ldflags "$(GO_LDFLAGS)" main.go healthz.go

run:
	go run -ldflags "$(GO_LDFLAGS)" main.go healthz.go
