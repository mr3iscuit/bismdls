BINARY_NAME=bisgols
BUILD_DIR=build/bin/

build:
	@mkdir -p ${BUILD_DIR}
	go build -o ${BUILD_DIR}/${BINARY_NAME} ./

.PHONY: build
