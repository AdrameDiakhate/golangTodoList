FROM golang:1.18-alpine AS gotodo

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . ./

COPY configs/.env .

RUN go build -o golangtodolist 

EXPOSE 9091

CMD [ "./golangtodolist" ]