# Build stage
# Name this stage as "builder"
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# No need to install migrate in the build stage because it runs migration in main.go file
# Install curl to download migrate
# RUN apk add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.18
WORKDIR /app
# Copy the executable file from the "builder" stage
COPY --from=builder /app/main .
# COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

# Publish docker image through port 8080
EXPOSE 8080 9090
CMD ["/app/main"]
# If `CMD` is used to provide default arguments for the `ENTRYPOINT` instruction, 
# both the `CMD` and `ENTRYPOINT` instructions should be specified with the JSON array format.
# In another word, it should be like this:
# ENTRYPOINT [ "/app/start.sh", "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]