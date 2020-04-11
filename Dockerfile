# First stage
FROM golang:1.13

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Make app directory
RUN mkdir /app

# Add app directory 
ADD . /app

# Add this go mod download command to pull in any dependencies
# RUN go mod download

# Create and change to the app directory.
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy local code to the container image.
ADD . .


# Build the binary.
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN go build -o cmd/main .

# Second stage
FROM alpine:latest

# Run the web service on container startup.
CMD ["/app"]

# Copy from first stage
COPY --from=0 /app/main /app