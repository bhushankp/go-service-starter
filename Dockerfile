FROM golang:1.25-alpine AS builder


WORKDIR /app

COPY go.mod ./

RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server .


#Runtime stage
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/server ./server

EXPOSE 8080

CMD ["./server"]