img_name := "end-of-life-cli"
owner    := "ludovanorden"
registry := "eol"

# Build the Docker container.
build:
    docker build . -t {{img_name}}

# Build the container specifically for testing.
build-test:
    docker compose build

# Run a command in the Docker container.
run *cmd:
    docker run --rm {{img_name}} {{cmd}}

# Run linter
lint: build-test
    docker compose run --rm test golangci-lint run /app/...

# Run unit tests in the Docker container.
test: build-test
    docker compose run --rm test
    docker compose down

cov: test
    cd src && go tool cover -html=../coverage/coverage.out

docker-builder := "eol-builder"

setup-buildx:
    docker buildx create --name {{docker-builder}} --driver docker-container --bootstrap

buildx tag="latest":
    docker buildx use {{docker-builder}}
    docker buildx build --platform linux/amd64,linux/arm64 -t {{owner}}/{{registry}}:{{tag}} --push .
