FROM golang:1.20-alpine AS builder

WORKDIR /app

# Kopiere und lade Abhängigkeiten
COPY go.mod go.sum ./
RUN go mod download

# Kopiere den Rest des Codes und baue das Binary
COPY . .
RUN go build -ldflags="-s -w" -o consumer-inventory .

# Verwende ein minimales Image für die Laufzeit
FROM alpine:latest
WORKDIR /root/
RUN apk --no-cache add ca-certificates

# Kopiere das gebaute Binary
COPY --from=builder /app/consumer-inventory .

CMD ["./consumer-inventory"]