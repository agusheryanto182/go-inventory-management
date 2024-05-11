FROM golang:1.22-alpine

COPY . /main_AgusHeryanto182

WORKDIR /main_AgusHeryanto182

RUN go mod tidy

RUN go build -o main_AgusHeryanto182 .

EXPOSE 8080

CMD ["/main_AgusHeryanto182/main_AgusHeryanto182"]

