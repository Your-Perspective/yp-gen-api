# Step 1: Build the Go application
FROM golang:1.23 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go app with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/yp-blog-api

# Step 2: Create a small image for running the application
FROM gcr.io/distroless/static-debian11

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the builder stage
COPY --from=builder /go/bin/yp-blog-api .

# Expose port 9090 to the outside world
EXPOSE 9090

# Command to run the executable
CMD ["./yp-blog-api"]
