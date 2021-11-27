FROM golang:alpine as builder
LABEL stage=builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY ./shippy-service-user ./shippy-service-user
WORKDIR /app/shippy-service-user

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-user *.go

#RUN CONTAINER
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shippy-service-user .

CMD ["./shippy-service-user"]
