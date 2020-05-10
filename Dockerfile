FROM golang:alpine
ADD . /go/src/HIldebrandoFerreira/apigo
RUN cd /go/src/HIldebrandoFerreira/apigo && go install
CMD ["/go/bin/apigo"]
EXPOSE 9090