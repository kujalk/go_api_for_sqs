FROM golang:latest


# Set the working directory
WORKDIR /app

# Copy the source code
COPY app/ .

# Install dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the Go application
CMD ["./main"]