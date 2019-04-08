FROM golang:1.12-alpine as build
MAINTAINER apt
WORKDIR $GOPATH/src/github.com/abhayprakashtiwari/estatebidding
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
CMD ["estatebidding" , "connectionUri=mongodb://estate-mongo:27017"]