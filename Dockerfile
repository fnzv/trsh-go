FROM golang:1.13

RUN mkdir -p /app

WORKDIR /app

RUN go get -u "gopkg.in/telegram-bot-api.v4"
RUN apt update -y
RUN apt install -y mtr dnsutils nmap net-tools

ADD . /app

RUN go build ./trsh.go

CMD ["./trsh"]
