FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY ./shippy-service-vessel ./shippy-service-vessel
WORKDIR /app/shippy-service-vessel  

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-vessel *.go

#RUN CONTAINER
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shippy-service-vessel .

CMD ["./shippy-service-vessel"]
