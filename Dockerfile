FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /goapiondocker

EXPOSE 8001

CMD [ "/goapiondocker" ]


#  docker build --tag goapiondocker .
#  docker image tag goapiondocker:latest goapiondocker:v1.0
#  docker image rm goapiondocker:v1.0