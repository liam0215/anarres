BUILD_DIR := build
DOCKER_DIR := docker

.PHONY: all build env kompose clean

all: build kompose

build:
	cd wiring && go run main.go -o ../$(BUILD_DIR) -w docker

kompose:
	cd $(BUILD_DIR) && \
	set -a && . ./.env && set +a && \
	cd $(DOCKER_DIR) && kompose convert -f docker-compose.yml && \
	mv docker-compose.yml docker-compose.yml.bak

run:
	cd $(BUILD_DIR)/$(DOCKER_DIR) && \
	kubectl apply -f .

clean:
	rm -rf $(BUILD_DIR)
