APP_NAME := server
BUILD_DIR := bin
MAIN_PKG := ./app

.PHONY: build run test vet clean docker-build docker-run

build:
	@mkdir -p $(BUILD_DIR)
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PKG)

run:
	go run $(MAIN_PKG)

test:
	go test -v ./app/tests/

vet:
	go vet ./app/...

clean:
	rm -rf $(BUILD_DIR)

docker-build:
	docker build -t $(APP_NAME) .

docker-run: docker-build
	docker run --rm -p 8080:8080 \
		-e JWT_SECRET=change-me \
		-e GIN_MODE=release \
		$(APP_NAME)
