#FROM --platform=linux/amd64 golang:1.20
FROM golang:1.20

WORKDIR /app
RUN mkdir /coverage

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

COPY ./src ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /eol
RUN mv /eol /usr/bin

CMD ["eol"]
