# EOL
Simple CLI application to check whether a tool's version has reached end-of-life (EOL).

EOL uses https://endoflife.date/ as a data source.

## Docker
Build the Docker container with `just build`. This will create a container named `end-of-life-cli`. No platform has been specified in the Dockerfile, so it will 
build for your machine's platform.

Run a command in the container with `just run eol ruby 2.7`.
