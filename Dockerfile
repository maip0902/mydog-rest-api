FROM golang:latest

RUN apt-get update && \
    apt-get -y install vim
RUN mkdir /go/src/mydog-api-app
WORKDIR /go/src/mydog-api-app
COPY . /go/src/mydog-api-app
#CMD ["./build/build"]