FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/gorilla/mux \
&&  go get -u github.com/go-sql-driver/mysql \
&&  go get -u github.com/DATA-DOG/go-sqlmock
#RUN go get -d -v ./...
#RUN go install -v ./...

CMD ["go", "run", "main.go"]