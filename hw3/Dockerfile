FROM golang:1.24.7-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /server ./cmd/main.go


FROM alpine:latest

COPY --from=builder /server /server

EXPOSE 8080

CMD ["/server"]
