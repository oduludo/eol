img_name := "end-of-life-cli"

build:
    docker build . -t {{img_name}}

run *cmd:
    docker run {{img_name}} {{cmd}}
