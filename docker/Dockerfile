FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /app/main

# Runtime Stage
FROM alpine:3.15
WORKDIR /bin

COPY --from=build /app/main /bin/main
COPY --from=build /app/config.yml /bin/config.yml
COPY --from=build /app/repository/mysql/migrations /bin/repository/mysql/migrations

EXPOSE 8787

CMD [ "/bin/main" ]