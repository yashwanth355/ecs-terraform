FROM golang:1.16
ENV GOFLAGS=-mod=readonly
ENV BTS_DIR=./cclLeapfrog
ENV APP_ENV production
ENV PORT 3000
ADD . $BTS_DIR

WORKDIR $BTS_DIR/cmd

RUN go build -v -o server
EXPOSE 8080
CMD ["./server"]
