FROM golang:alpine
RUN mkdir -p /go/src && mkdir -p /go/bin && mkdir -p /go/pkg

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH 

RUN mkdir -p $GOPATH/src/go-examples/addressbook 
ADD . $GOPATH/src/go-examples/addressbook

WORKDIR $GOPATH/src/go-examples/addressbook 
RUN go build -o main .
CMD ["/go/src/go-examples/addressbook/main"]