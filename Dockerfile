# Dockerfile
FROM golang:alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum .env ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./
COPY cmd/ ./cmd/
COPY mocks/ ./mocks/
COPY pkg/ ./pkg/

# Build
RUN go build -o /hands-on-go ./cmd/web/

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/hands-on-go"]