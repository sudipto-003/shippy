FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY ./shippy-service-consignment ./shippy-service-consignment
ADD ./shippy-service-vessel ./shippy-service-vessel
WORKDIR /app/shippy-service-consignment  

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-consignment *.go

#RUN CONTAINER
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shippy-service-consignment .

CMD ["./shippy-service-consignment"]
