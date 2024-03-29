FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod go.sum /app/
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE  8080

ENTRYPOINT [ "./main" ]