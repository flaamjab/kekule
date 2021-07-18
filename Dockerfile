FROM golang
WORKDIR /kekule
COPY go.mod .
COPY go.sum .
COPY internal internal
COPY hack hack
COPY cmd cmd
RUN go mod download
CMD go run cmd/kekule/server.go $ADDRESS
