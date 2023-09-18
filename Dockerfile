FROM golang:1.21.1 AS build

WORKDIR /usr/www
ENV GOOS=linux GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./api .

FROM alpine:3.17 AS prod

RUN apk update && \
  apk add --no-cache ca-certificates openssl dumb-init && \
  rm -rf /var/cache/apk/*

WORKDIR /usr/www

COPY --from=build /usr/www/api ./
COPY ./start.sh ./

RUN chmod +x ./api && chmod +x ./start.sh

RUN adduser -D go

RUN chown go:go /usr/www/

EXPOSE 5005

USER go

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/bin/sh", "-c", "./start.sh"]
