vendor:
	export GO111MODULE=on
	go mod vendor
	export GO111MODULE=off

build:
	go vet
	go build -o bin/application application.go

build-npm:
	cd web/ && npm install && npm run build -- --go

vendor-update:
	export GO111MODULE=on
	go build
	export GO111MODULE=off

build-web:
	cd web/ && npm run build -- --go
