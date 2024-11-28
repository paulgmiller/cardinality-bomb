FROM golang:1.23 AS builder
ENV PORT 8080
EXPOSE 8080

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app-binary

FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=builder /build/app-binary . 
CMD ["/app/app-binary"]
