FROM golang:1.17.6-alpine3.15

WORKDIR /deckofcards

RUN export GOROOT="/usr/local/go"
#RUN go env -w GOPATH=$HOME/go


COPY go.mod ./
COPY go.sum ./
RUN go mod download


COPY *.go ./
COPY app/ ./app

#RUN echo $GOPATH


RUN go build -o deckofcards-go .

EXPOSE 8080

CMD [ "./deckofcards-go" ]
