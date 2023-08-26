FROM golang:1.19-alpine as builder 
# Create app directory
WORKDIR /app
# Copy go mod and sum files
COPY . /app

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x /app/brokerApp

# Build the Go app
RUN go build -o broker-service ./cmd/api

FROM alpine:latest  

RUN mkdir /app



COPY --from=builder /app/brokerApp /app

# Command to run the executable
CMD ["/app/brokerApp"]