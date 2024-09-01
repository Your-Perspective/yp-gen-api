# Step 1: Build the Go application
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire source code to the container
COPY . .

# Enable CGO and build the Go app
RUN CGO_ENABLED=1 GOOS=linux go build -o /go/bin/yp-blog-api ./

# Step 2: Use a compatible base image for running the application
FROM ubuntu:22.04

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the builder stage
COPY --from=builder /go/bin/yp-blog-api .
COPY .env.fly ./
# Copy the SQLite database file from the local machine to the container
COPY test.db /root/test.db
ARG APP_ENV=fly
ENV APP_ENV=${APP_ENV}
# Expose port 9090 to the outside world
EXPOSE 9090

# Command to run the executable
CMD ["./yp-blog-api"]