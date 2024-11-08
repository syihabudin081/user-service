# Use the official Golang image as the base image
FROM golang:1.23

# Set the working directory to /app
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Expose the port that the application will run on
EXPOSE 8080

# Set the command to run the application
CMD ["./main"]