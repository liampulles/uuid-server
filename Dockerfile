FROM golang:1.14-alpine
RUN apk update && apk add --no-cache git

ENV GOPATH /go
ENV GOBIN /go/bin
ENV CGO_ENABLED 0
WORKDIR /app
RUN mkdir -p /go/bin

COPY ./go.mod .
RUN go mod download
RUN go mod verify

COPY . .
RUN go build ./...
RUN go install ./...

FROM scratch
COPY --from=0 "/go/bin/uuid-server" "/go/bin/uuid-server"
ENTRYPOINT ["/go/bin/uuid-server"]