FROM golang:1.12-alpine as build
MAINTAINER apt
#RUN mkdir /app
#ADD . /app/
#WORKDIR /app
#WORKDIR /go/src/app
WORKDIR $GOPATH/src/github.com/abhayprakashtiwari/estatebidding
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
#RUN wget https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
#RUN go get -u github.com/golang/dep/cmd/dep
#RUN dep ensure
#RUN go build -o app
#FROM alpine:3.7
#COPY --from=build /go/src/app/app /usr/local/bin/app
#CMD ["/usr/local/bin/app"]
CMD ["estatebidding" , "connectionUri=mongodb://estate-mongo:27017"]