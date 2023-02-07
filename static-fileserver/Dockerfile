FROM golang:latest

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download && go mod verify

COPY . /app/

RUN go build -v -o /usr/local/bin/static-fileserver ./...
EXPOSE 8080

CMD ["/usr/local/bin/static-fileserver"]
