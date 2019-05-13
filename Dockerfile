FROM golang

RUN go get github.com/gorilla/mux

RUN go get golang.org/x/crypto/ssh

WORKDIR /go/src/github.com/heaptracetechnology/microservice-ssh

ADD . /go/src/github.com/heaptracetechnology/microservice-ssh

RUN go install github.com/heaptracetechnology/microservice-ssh

ENTRYPOINT microservice-ssh

EXPOSE 3000