# Build image
FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . .
ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o ./out/app .

# Runtime image
FROM alpine:3.20
COPY --from=build /app/out/app /usr/local/bin/app
EXPOSE 443
CMD ["/usr/local/bin/app"]