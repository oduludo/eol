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
    docker run {{img_name}} {{cmd}}

# Publish the Docker container to Docker Hub.
# When working on an ARM machine, make sure to build for the correct target (which can be specified in the Dockerfile).
docker-publish: build
    docker tag {{img_name}} {{owner}}/{{registry}}:latest
    docker push {{owner}}/{{registry}}:latest

# Run linter
lint: build-test
    just run golangci-lint run /app/...

# Run unit tests in the Docker container.
test: build-test
    docker compose run --rm test

cov: test
    cd src && go tool cover -html=../coverage/coverage.out
