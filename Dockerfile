# Use official Golang image version 1.19.1 as the base image
FROM golang:1.19.1

# Set the Current Working Directory inside the container
WORKDIR /app

# I am not installing any third-party packages in this Dockerfile
# as the Go application does not depend on any external dependencies.

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]

