FROM golang:1.18 AS gotodo

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy

COPY . ./

RUN go build -o /golangtodolist

FROM alpine:3.15.4

WORKDIR /app

COPY --from=gotodo golangtodolist .

EXPOSE 9090

CMD [ "./golangtodolist" ]