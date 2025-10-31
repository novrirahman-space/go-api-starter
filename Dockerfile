# Build stage
FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/server

# Run stage (distroless)
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/server .
ENV APP_ENV=production
ENV HTTP_ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/app/server"]
