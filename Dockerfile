# Stage build: gunakan image resmi Go
FROM golang:1.23 AS builder

WORKDIR /app

# Salin file modul dan unduh dependency
COPY go.mod go.sum ./
RUN go mod download

# Salin seluruh source code
COPY . .

# Build aplikasi
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Stage runtime: gunakan image yang ringan (misalnya Alpine)
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Salin binary dari stage builder
COPY --from=builder /app/app .

# Expose port 8080
EXPOSE 8080

# Jalankan aplikasi
CMD ["./app"]
