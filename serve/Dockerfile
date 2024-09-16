FROM golang:1.18-alpine

WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY serve/go.mod serve/go.sum ./
COPY shared/go.mod shared/go.sum ./shared/
RUN go mod download

# Copy the source code
COPY serve/ ./
COPY shared/ ./shared/

# Build the service
RUN go build -o serve .

EXPOSE 8000
CMD ["./maintain"]
