FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Copy the binary file from the build context to the container
COPY ./main .

# Copy static files
COPY ./web ./web

# Expose the port on which the application is running
EXPOSE 8080

# Run the binary
CMD ["./main"]
