FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/sjsnider/nds

RUN go install github.com/sjsnider/nds

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/nds

# Document that the service listens on port 8080.
EXPOSE 8080
