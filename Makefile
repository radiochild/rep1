APP_NAME=rep1
APP_DESC="produces\ reporting\ data\ that\ is\ either\ published\ to\ S3\ or\ saved\ to\ an\ accessible\ file\ system"
REPO_ORG=github.com/radiochild
MOD=${REPO_ORG}/${APP_NAME}

# -----------------------------------------------------------------------------------
#  The file name VERSION should contain the app version without any end-of-line
#  characteres.  The recommended format is vX.Y.Z (following Semantic Versioning 2.0
#  conventions.  Example: v1.0.4
# -----------------------------------------------------------------------------------
VERSION_FILE := VERSION

# -----------------------------------------------------------------------------------
# Inject Name, Version and Build info into the binary
# -----------------------------------------------------------------------------------
BUILD_VERSION=$(shell cat ${VERSION_FILE})
BUILD_REVISION=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date '+%Y-%b-%d')
BUILD_TIME=$(shell date '+%H:%M:%S %Z')
APP_LDFLAGS=-ldflags "\
			-X 'main.AppName=${APP_NAME}' \
			-X 'main.AppDesc=${APP_DESC}' \
			-X 'main.Version=${BUILD_VERSION}' \
			-X 'main.Build=${BUILD_REVISION}' \
			-X 'main.BuildDate=${BUILD_DATE}' \
			-X 'main.BuildTime=${BUILD_TIME}'"

# -----------------------------------------------------------------------------------
# Publishing Docker images
# -----------------------------------------------------------------------------------
IMAGE_NAME="${APP_NAME}:${BUILD_VERSION}"
DOCKERIZED_NAME="${APP_NAME}-docker"

# -----------------------------------------------------------------------------------
# Build Tools
# -----------------------------------------------------------------------------------
LINTER=golangci-lint
#LINTER_FLAGS=
LINTER_FLAGS=-v

OUT_DIR=build
.PHONY: clean

clean:
	@rm -rf ${OUT_DIR}
	@mkdir -p ${OUT_DIR}

init:	clean
	@go mod init ${MOD}
	@go get

lint:	clean
	# ${LINTER} run ${LINTER_FLAGS} --enable=gosec --enable=revive --skip-dirs build

test:	lint
	# go test -covermode=count -coverprofile=${OUT_DIR}/pkgc.out ./...
	# go tool cover -func=${OUT_DIR}/pkgc.out

build:	test
	go build ${APP_LDFLAGS} -a -o ${OUT_DIR}/${APP_NAME}

image:	build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build ${APP_LDFLAGS} -a -o ${OUT_DIR}/${DOCKERIZED_NAME}
	docker build -t ${IMAGE_NAME} .


