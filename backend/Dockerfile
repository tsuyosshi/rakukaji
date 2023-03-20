FROM golang:1.19.1-alpine

RUN mkdir /go/src/rakukaji

WORKDIR /go/src/rakukaji

COPY ./backend/src ./

COPY ./backend/run.sh ./run.sh
RUN chmod +x ./run.sh
CMD ["./run.sh"]