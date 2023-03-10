FROM golang:1.20 AS builder

# Download dependencies in a separate layer to promote caching.
WORKDIR /usr/src/example-app
COPY go.mod .
RUN go mod download && \
	go mod verify

# Compile the application.
COPY . .
RUN go build -v -o /usr/local/bin/example-app ./...

# Final container image.
FROM ubuntu:jammy
COPY --from=builder /usr/local/bin/example-app /usr/local/bin/example-app
CMD ["example-app"]
