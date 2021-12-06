FROM golang:1.16-alpine3.13
RUN apk add --no-cache git gcc
RUN mkdir /app
WORKDIR /app
#WORKDIRを動かさないと「$GOPATH/go.mod exists but should not」と怒られる
COPY ./APIserver .
RUN go mod download \
    && go get github.com/pilu/fresh
EXPOSE 8080
ENTRYPOINT ["fresh"]
#変更に追従してビルドしてくれるやつ
