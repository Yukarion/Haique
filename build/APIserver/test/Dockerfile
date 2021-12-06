FROM golang:1.16-alpine3.13
RUN apk add --no-cache git gcc musl-dev
RUN mkdir /app
WORKDIR /app
COPY ./APIserver .
RUN go mod download
EXPOSE 8080
ENTRYPOINT ["go","test"]
