FROM ubuntu:22.04

RUN apt-get update && apt-get install -y python3 python3-pip

WORKDIR /usr/src/app

COPY ./dist/ ./

ENV GREETINGS_PATH="greetings.json"

CMD ["./pyhello"]
