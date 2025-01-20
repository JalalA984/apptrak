# Use the official Golang image from Docker Hub
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files (if you're using Go modules)
COPY go.mod go.sum ./

# Download and cache dependencies (if any)
RUN go mod download

# Copy the rest of the application code into the container
COPY . .

# Build the Go application. The output is placed in /app/cmd/apptrak.
RUN go build -o /app/cmd/apptrak/main ./cmd/apptrak

# Expose the port your application will run on (e.g., port 8080)
EXPOSE 5000

# Command to run the application
CMD ["./cmd/apptrak/main"]
