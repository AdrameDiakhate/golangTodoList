FROM golang:1.18-alpine AS gotodo

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . ./

RUN go build -o golangtodolist 

FROM alpine

WORKDIR /app

COPY --from=gotodo /app/golangtodolist .

COPY --from=gotodo /app/.env .

EXPOSE 9091


CMD [ "./golangtodolist" ]