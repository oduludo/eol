img_name := "end-of-life-cli"
owner    := "ludovanorden"
registry := "eol"

# Build the Docker container.
build:
    docker build . -t {{img_name}}

# Run a command in the Docker container.
run *cmd:
    docker run {{img_name}} {{cmd}}

# Publish the Docker container to Docker Hub.
# When working on an ARM machine, make sure to build for the correct target (which can be specified in the Dockerfile).
docker-publish: build
    docker tag {{img_name}} {{owner}}/{{registry}}:latest
    docker push {{owner}}/{{registry}}:latest

# Run unit tests in the Docker container.
test: build
    docker compose run --rm test

cov: test
    cd src && go tool cover -html=../coverage/coverage.out
