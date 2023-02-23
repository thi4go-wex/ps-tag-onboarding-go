FROM golang:alpine

RUN apk update && apk add --no-cache go

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

COPY . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]do
