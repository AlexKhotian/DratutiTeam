FROM golang:1.8
USER 1001
WORKDIR /go/src/DratutiTeam/WebBackend
COPY . /go/src/DratutiTeam/WebBackend
RUN go test 
RUN go install 
EXPOSE 7777
ENTRYPOINT [ "/go/bin/WebBackend"]
