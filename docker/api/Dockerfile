FROM golang:1.19-alpine
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /go/src/app
WORKDIR /go/src/app

CMD ["go", "run", "main.go"]