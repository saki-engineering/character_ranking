FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/gorilla/mux \
&&  go get github.com/gomodule/redigo/redis \
&&  go get github.com/google/uuid
#RUN go get -d -v ./...
#RUN go install -v ./...

CMD ["go", "run", "main.go"]