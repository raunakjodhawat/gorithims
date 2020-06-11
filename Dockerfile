FROM golang:1.14
# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/raunakjodhawat/gorithims

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Run tests
RUN go test -v -race ./... --cover

# Run the executable
CMD ["go", "run", "src/dockerExecution.go"]