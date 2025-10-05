# ---------- Build Stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git & bash (needed for go modules and scripts)
RUN apk add --no-cache git bash

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o main ./cmd/main.go

# ---------- Run Stage ----------
FROM golang:1.25-alpine

WORKDIR /app

# Install bash & curl
RUN apk add --no-cache bash curl

# Copy binary, env file, and Air config
COPY --from=builder /app/main .
COPY .env .env
COPY .air.toml .air.toml

# Install Air for live reload
RUN curl -fSL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh \
    && mv ./bin/air /usr/local/bin/air

EXPOSE 8080

# Default CMD: use Air if .air.toml exists, else fallback to static binary
CMD ["sh", "-c", "export PATH=$PATH:/usr/local/bin && if [ -f .air.toml ]; then air -c .air.toml; else ./main; fi"]
