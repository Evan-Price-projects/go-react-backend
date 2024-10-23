# Use the official Golang image as the base image
FROM golang:1.23-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Use a smaller base image for the final image
FROM alpine:latest

# Copy the executable from the build stage
COPY --from=build /app/myapp /myapp

# Expose the port that your application listens on (if applicable)
EXPOSE 8080

# Set the command to run the application
CMD ["/myapp"]