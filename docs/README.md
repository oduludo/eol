# EOL
Simple CLI application to check whether a tool's version has reached end-of-life (EOL).
EOL is designed primarily for usage in CI/CD pipelines, optionally giving an exit code if the EOL has passed.
Docker images are available to be used in jobs, but it can be used directly on your workstation as well.

EOL uses https://endoflife.date/ as a datasource. This datasource will be referred to as the 'root' source.

Links:
- Docker images can be found on [Docker Hub](https://hub.docker.com/r/ludovanorden/eol).
- Getting started guide is found [here](./getting_started.md).
- Read more on the [usage](./usage.md) of EOL.
