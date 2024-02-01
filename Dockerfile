FROM golang:alpine as builder
LABEL authors="KotopheiOP"

WORKDIR /build
ADD go.mod .
ADD go.sum .
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/todo-list ./cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/todo-list /app/todo-list

EXPOSE 8080
CMD ["./todo-list"]