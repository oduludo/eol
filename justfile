img_name := "end-of-life-cli"
owner    := "ludovanorden"
registry := "eol"

build:
    docker build . -t {{img_name}}

run *cmd:
    docker run {{img_name}} {{cmd}}

docker-publish: build
    docker tag {{img_name}} {{owner}}/{{registry}}:latest
    docker push {{owner}}/{{registry}}:latest
