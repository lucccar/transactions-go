FROM golang:latest
USER root

COPY . $GOPATH/src/transactions
WORKDIR $GOPATH/src/transactions/app

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -o ./bin/ -v ./...
RUN chmod -R 777 ./bin
RUN chmod -R 777 ./bin/src

EXPOSE 8080

CMD ["./bin/src"]