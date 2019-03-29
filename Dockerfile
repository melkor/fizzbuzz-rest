# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang v1.11 base image
FROM golang:1.11

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/melkor/fizzbuzz-rest

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Install the package
RUN make 

# This container exposes port 8080 to the outside world
EXPOSE 8000

# Run the executable
CMD ["fizzbuzz-rest", "-d", "-v"]
