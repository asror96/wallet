FROM golang:1.23

RUN go version
ENV GOPATH=/

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod download
RUN go build -o wallet ./cmd/main.go
RUN go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
CMD ["./wallet"]
