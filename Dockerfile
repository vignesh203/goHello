FROM ubuntu:18.04
# Install Stable Go
FROM golang

ADD . /opt/goHello
ENV PATH /usr/local/go/bin:/usr/local/bin:$PATH
#ENV GOPATH /opt/goHello
#ENV GOBIN $GOPATH/bin

# Install goHello
WORKDIR /opt/goHello
RUN go install goHello
