# Build with golang base image
FROM golang:1.14 as builder

ENV GO111MODULE=on

# Meta info
LABEL Company="Momolab"
LABEL Project="Fuselink backend"

# Set the current working directory inside the container
WORKDIR /fuselink

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go executable from source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fuselink .

# Start a new stage from with alpine linux
FROM alpine:3.10.4

COPY --from=builder /fuselink/fuselink .

# Expose port 8080
EXPOSE 3000

# Run the fuse executable
ENTRYPOINT ["./fuselink"]
