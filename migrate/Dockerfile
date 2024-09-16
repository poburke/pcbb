FROM golang:1.18-alpine

WORKDIR /app

# Copy the go.mod and go.sum files first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Run migrations manually
CMD ["go", "run", "main.go"]
