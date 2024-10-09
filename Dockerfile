# Use a lightweight Go base image
FROM golang:1.23-alpine

# Install bash (or sh) to allow exec access inside the container
RUN apk add --no-cache bash

# Set working directory inside the container
WORKDIR /app

# Copy the Go application into the container
COPY . /app

# Build the Go application
RUN go build -o Diskreader .

# Define the command to run the application
CMD ["/app/Diskreader"]