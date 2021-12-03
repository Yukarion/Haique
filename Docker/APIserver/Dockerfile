FROM golang:1.16-alpine3.13 as build-env
RUN apk add --no-cache git gcc
RUN mkdir /app
WORKDIR /app
COPY ./APIserver .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o openapi
FROM alpine:3.13 as prod
COPY --from=build-env /app/openapi .
EXPOSE 8080/tcp
USER 1001
ENTRYPOINT ["./openapi"]

FROM golang:1.16-alpine3.13 as dev
RUN apk add --no-cache git gcc\
    && go get github.com/pilu/fresh 
RUN mkdir /app
WORKDIR /app
#WORKDIRを動かさないと「$GOPATH/go.mod exists but should not」と怒られる
COPY ./APIserver .
EXPOSE 8080
ENTRYPOINT ["fresh"]
#変更に追従してビルドしてくれるやつ