# Use Alpine Linux as the base image
FROM alpine:edge

# Set environment variables
ENV GO_VERSION=1.24.0

# Install dependencies
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
    apk update && \
    apk add --no-cache \
    go=${GO_VERSION}-r0    

# Set the working directory
WORKDIR /app

# Copy the Go project files into the container
COPY . .

# Download Go dependencies
RUN go mod download

# Build the Go project
RUN go build -o main main.go

# Expose the port your Go application will run on
EXPOSE 8011

# Command to run your Go application
CMD ["/app/main"]