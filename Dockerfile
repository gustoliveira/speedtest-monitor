FROM golang:1.22.2

RUN \
  apt update && \
  apt install curl -y && \
  curl -s https://packagecloud.io/install/repositories/ookla/speedtest-cli/script.deb.sh | bash && \
  apt install speedtest

RUN mkdir database

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o speedtest

CMD ["./speedtest"]
