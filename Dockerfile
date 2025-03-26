# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o tree-visualization

# Final stage
FROM alpine:latest
RUN apk add --no-cache graphviz
WORKDIR /app
COPY --from=builder /app/tree-visualization .
EXPOSE 8080
ENTRYPOINT ["./tree-visualization"]