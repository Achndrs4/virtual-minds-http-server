# Start from a minimal golang image
FROM golang:alpine AS build

WORKDIR /app

# Copy the local package files to the container's workspace
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
# Build the Go app
RUN go build -o server

# Start from a smaller base image
FROM alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=build /app/server .
COPY --from=build /app/config.yaml .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./server"]