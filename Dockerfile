FROM golang:1.22.2

WORKDIR /app

RUN \
  apt update && \
  apt install curl -y && \
  curl -s https://packagecloud.io/install/repositories/ookla/speedtest-cli/script.deb.sh | bash && \
  apt install speedtest -y && \
  mkdir database

COPY go.mod go.sum speedtest.go /app/

RUN go mod download

RUN go build -o speedtest

CMD ["./speedtest"]
