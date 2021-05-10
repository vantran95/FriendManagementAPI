# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="VanTran <vantx95@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /api

# Copy go mod and sum files
COPY api/go.mod api/go.sum ./
COPY api/.env ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY api/cmd/serverd .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /api/cmd/serverd/main .
COPY --from=builder /api/.env .

# Expose port 8083 to the outside world
EXPOSE 8082

#Command to run the executable
CMD ["./main"]