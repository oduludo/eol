#FROM --platform=linux/amd64 golang:1.20
FROM --platform=$BUILDPLATFORM golang:1.20 as base

WORKDIR /app
RUN mkdir /coverage

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

# Source files
COPY ./src ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /eol
RUN mv /eol /usr/bin

CMD ["eol"]

FROM base as test

# Mimic having some datasource file somewhere in the file system to process
RUN mkdir /home/files
COPY ./docs/example_datasource.json /home/files/example_datasource_1.json
COPY ./docs/example_datasource.json /home/files/example_datasource_2.json
COPY ./docs/example_datasource.json /home/files/example_datasource_3.json
COPY ./docs/example_datasource.json /home/files/example_datasource_4.json
