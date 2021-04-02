FROM golang:1.15
ENV APP_HOME go/src/github.com/syned13/clinics-api

RUN mkdir -p $APP_HOME
ADD . $APP_HOME
WORKDIR $APP_HOME

RUN go mod download
RUN go build -o main ./cmd/main.go

RUN chmod +x ./main

CMD ["./main"]