FROM golang:1.11.1

ENV GO111MODULE=on
EXPOSE 80

COPY .. /home

RUN cd /home \
  &&  go build -o svc-demo .

CMD /home/svc-demo --port=80