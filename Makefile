APP=nexus-enedis-m23-measure-lambda

.PHONY: tidy run build package clean

tidy:
	go mod tidy

run:
	bash ./run-local.sh

build:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -tags lambda -ldflags="-s -w" -o bootstrap ./main.go

package: build
	zip -q function.zip bootstrap

clean:
	rm -f bootstrap function.zip
