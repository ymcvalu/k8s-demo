FROM golang:1.11

EXPOSE 80

WORKDIR /home

COPY . .

RUN go build -o ping-pong

CMD ./ping-pong --port=80