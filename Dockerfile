FROM ubuntu:18.04
# Install Stable Go
FROM golang

ADD . /opt/goHello
ENV PATH /usr/local/go/bin:/usr/local/bin:$PATH

# Install goHello
WORKDIR /opt/goHello
RUN go install goHello
ENTRYPOINT goHello
EXPOSE 80
