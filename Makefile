VERSION := 1.0.0
COMMIT_HASH := ` git rev-parse --short HEAD `

dep:
	which godep || go get github.com/tools/godep

test: build
	go test -race -v ./...

vet:
	go list ./... | grep -v "./vendor*" | xargs go vet

fmt:
	find . -type f -name "*.go" | grep -v "./vendor*" | xargs gofmt -s -w

build: dep fmt vet
	godep go install -v ./...
	godep go build -v -ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH)" -o hostipreflection github.com/andyxning/host_ip_reflection/cmd/hostipreflection

clean:
	rm hostipreflection

run:
	./hostipreflection --logtostderr=true

.PHONY: fmt test dep build clean run vet
