FROM golang:1.11
EXPOSE 80
RUN touch /go/src/main.log
COPY ./bin/hello-server /usr/local/bin/
CMD ["hello-server"]
