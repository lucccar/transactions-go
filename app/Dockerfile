FROM golang:latest
USER root

COPY . $GOPATH/src/transactions
WORKDIR $GOPATH/src/transactions/app

ENV EXCHANGE_BASE_URL="https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?"
ENV DBHOST="db"
ENV DBPORT="5432"
ENV DBUSER="postgres"
ENV DBPASSWORD "carvalho"
ENV DBNAME="exchanegdb"

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -o ./bin/ -v ./...
RUN chmod -R 777 ./bin
RUN chmod -R 777 ./bin/src

EXPOSE 8080

CMD ["./bin/src"]