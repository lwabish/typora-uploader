version=v1.0.0

install:
	go install -ldflags " \
	-X main.version=$(version) \
    " cmd/tqu.go
