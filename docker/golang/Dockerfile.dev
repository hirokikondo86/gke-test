FROM golang:1.13.4

RUN mkdir /go/src/server
WORKDIR /go/src/server
COPY ./server .

RUN go get -u github.com/oxequa/realize
RUN go get -u github.com/golang/dep/cmd/dep

EXPOSE 9000