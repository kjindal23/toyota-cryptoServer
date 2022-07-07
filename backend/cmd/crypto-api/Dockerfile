FROM ubuntu:20.04

WORKDIR /usr/local/service

RUN apt-get update && apt-get install -y ca-certificates curl

COPY ./main .

ENTRYPOINT ["./main"]
EXPOSE 9090
EXPOSE 8080
