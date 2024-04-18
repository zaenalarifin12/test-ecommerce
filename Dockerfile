# Use a smaller base image for optimization
FROM golang:alpine3.19

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Download module dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8000 to allow communication to/from the container
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
