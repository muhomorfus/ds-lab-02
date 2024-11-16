FROM golang:1.23.2

COPY . /build
WORKDIR /build

RUN go mod tidy
RUN go build -o /opt/reservation /build/reservation/cmd/service/main.go

ENTRYPOINT ["/opt/service"]