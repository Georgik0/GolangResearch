FROM golang:1.17.2-alpine

WORKDIR /app

COPY go.mod ./
RUN go get github.com/lib/pq
RUN go get github.com/jmoiron/sqlx
RUN go mod download


COPY chat_avito.go ./

RUN go build -o /docker-gs-ping

EXPOSE 9000

CMD [ "/docker-gs-ping" ]