FROM golang:1.15 AS builder

WORKDIR /build

COPY . .

RUN go build ./cmd/main.go

FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install postgresql-12 -y

USER postgres

COPY ./init.sql .

RUN service postgresql start && \
    psql -c "CREATE USER statuser WITH superuser login password 'statuser';" && \
    createdb -O statuser statdb && \
    psql -c "ALTER ROLE statuser WITH PASSWORD 'statuser';" && \
    psql -d statdb < ./init.sql && \
    service postgresql stop

VOLUME ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

USER root

WORKDIR /proxy
COPY --from=builder /build/main .

COPY . .

EXPOSE 8080
EXPOSE 8000
EXPOSE 5432

ENV DB_USER=statuser
ENV DB_NAME=statdb

CMD service postgresql start && ./main
