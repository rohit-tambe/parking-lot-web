# stage 1 build golang binary
FROM golang:1.13 as build

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd

# stage 2 pack
FROM alpine:latest
COPY --from=build /app/main /

# add web folder
RUN mkdir web
COPY web /web

EXPOSE 8081

ENTRYPOINT ["/main"]

CMD [ "./setup" ]