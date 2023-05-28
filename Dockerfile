#FROM --platform=linux/amd64 golang:1.20
FROM --platform=$BUILDPLATFORM golang:1.20

WORKDIR /app
RUN mkdir /coverage

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

COPY ./src ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /eol
RUN mv /eol /usr/bin

CMD ["eol"]
