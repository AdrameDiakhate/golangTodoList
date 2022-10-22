FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY * ./

RUN go get .

RUN go mod tidy

RUN go build -o /golangtodolist

EXPOSE 9090

CMD [ "/golangtodolist" ]