FROM golang
ADD . /go/src/github.com/nebbers1111/webserver
RUN go install github.com/nebbers1111/webserver
ENTRYPOINT /go/bin/webserver
EXPOSE 8080