# Build image
FROM golang:1.16-alpine AS build
WORKDIR /app
COPY go-server/go.mod go-server/go.sum .
RUN go mod download
COPY go-server/. .
RUN go mod download
RUN go build -o ./out/app .

# Runtime image
FROM alpine:3.13
COPY --from=build /app/out/app /usr/local/bin/app
EXPOSE 8080
CMD ["/usr/local/bin/app"]