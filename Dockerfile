# Use the official Golang 1.22.4 image as a base
FROM golang:1.22.4

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's working directory
COPY . .

# Build the Go app
RUN go build -o backtester .

# Command to run the executable
CMD ["./backtester"]