FROM golang:1.16-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o goapiondocker ./cmd/app

# RUN chmod +x /app/brokerApp

# build a tiny docker image
# FROM alpine:latest

# RUN mkdir /app

# COPY --from=builder /app /app

EXPOSE 8001 8001

CMD [ "/app/goapiondocker" ]


#  docker build --tag goapiondocker .
#  docker start dazzling_chandrasekhar
#  docker stop dazzling_chandrasekhar
#  docker image tag goapiondocker:latest goapiondocker:v1.0
#  docker image rm goapiondocker:v1.0