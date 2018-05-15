
FROM golang

ADD . /go/src/room-one

RUN go install room-one
ENTRYPOINT /go/bin/room-one

EXPOSE 3000

CMD [ "encryption-service" ]