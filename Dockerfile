FROM golang:1.25.7

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o featurepilot ./cmd/api

# install migrate cli
RUN apt-get update && apt-get install -y wget

RUN wget -O /tmp/migrate.tar.gz \
https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz && \
tar -xzf /tmp/migrate.tar.gz -C /usr/local/bin && \
chmod +x /usr/local/bin/migrate

COPY migrations ./migrations

COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]