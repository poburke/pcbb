FROM golang:1.18-alpine

WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY maintain/go.mod maintain/go.sum ./
COPY shared/go.mod shared/go.sum ./shared/
RUN go mod download

# Copy the source code
COPY maintain/ ./
COPY shared/ ./shared/

# Build the service
RUN go build -o maintain .

EXPOSE 8080
CMD ["./maintain"]

