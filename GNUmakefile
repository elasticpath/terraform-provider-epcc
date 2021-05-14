NAME=epcc
HOSTNAME=elasticpath.com
NAMESPACE=elasticpath
BINARY=terraform-provider-${NAME}
VERSION=0.0.1

OS :=
ARCH :=
ifeq ($(OS),Windows_NT)
	OS = windows
	ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
		ARCH = amd64
	endif
	ifeq ($(PROCESSOR_ARCHITECTURE),x86)
		ARCH = 386
	endif
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		OS = linux
	endif
	ifeq ($(UNAME_S),Darwin)
		OS = darwin
	endif

	UNAME_M := $(shell uname -m)
	ifeq ($(UNAME_M),x86_64)
		ARCH=amd64
	endif
	ifneq ($(filter %86,$(UNAME_M)),)
		ARCH=386
	endif
	ifneq ($(filter arm%,$(UNAMEMP)),)
		ARCH=arm
	endif

endif
OS_ARCH=${OS}_${ARCH}

default: install

build:
	go build -o ./bin/${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64


install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv bin/${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

clean:
	rm -rf bin || true
	rm -rf ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME} || true

# Run acceptance tests
.PHONY: testacc
testacc:
	(\
		set -o allexport &&	[[ -f .env ]] && source ./.env && set +o allexport && \
		TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m \
	)

.PHONY: example
example: install
	(\
		set -o allexport &&	[[ -f .env ]] && source ./.env && set +o allexport && \
		pushd $(EXAMPLE) && \
		(rm .terraform.lock.hcl || true) && \
		terraform init && \
		terraform $(ACTION) && \
		popd \
	)

.PHONY: example
resource: install
	(\
		set -o allexport &&	[[ -f .env ]] && source ./.env && set +o allexport && \
		pushd examples/resources/$(TYPE)_resource && \
		(rm .terraform.lock.hcl || true) && \
		terraform init && \
		terraform $(ACTION) && \
		popd \
	)

.PHONY: example
data-source: install
	(\
		set -o allexport &&	[[ -f .env ]] && source ./.env && set +o allexport && \
		pushd examples/data-sources/$(TYPE)_data_source && \
		(rm .terraform.lock.hcl || true) && \
		terraform init && \
		terraform $(ACTION) && \
		popd \
	)


.PHONY: docs
docs: install
	(\
		set -o allexport &&	[[ -f .env ]] && source .env && set +o allexport && \
		(rm .terraform.lock.hcl || true) && \
		terraform init && \
		go generate \
	)

