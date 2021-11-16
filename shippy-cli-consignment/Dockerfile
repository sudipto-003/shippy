FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

#COPY ALL LOCAL DEPENDENCIES TO BUILD AND TEST LOCALLY
#USING: docker build -f ./shippy-cli-consignment/Dockerfile . -t shippy-cli-consignment:latest
#BUILD THE IMAGE FROM shippy/ directory
COPY ./shippy-cli-consignment ./shippy-cli-consignment
ADD ./shippy-service-consignment ./shippy-service-consignment
ADD ./shippy-service-vessel ./shippy-service-vessel
WORKDIR /app/shippy-cli-consignment

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-cli-consignment

#RUN CONTAINER
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shippy-cli-consignment/consignment.json .
COPY --from=builder /app/shippy-cli-consignment/shippy-cli-consignment .

CMD ["./shippy-cli-consignment"]
