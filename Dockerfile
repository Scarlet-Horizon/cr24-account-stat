FROM golang:1.19.3-alpine3.16

WORKDIR /api

COPY . .

RUN go mod init main && go mod tidy

RUN go mod download

RUN CGO_ENABLED=0 go build -o account-stat main.go

ENV GIN_MODE=release

CMD ["./account-stat"]
