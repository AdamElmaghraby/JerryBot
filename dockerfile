FROM golang:1.22

WORKDIR /app

COPY go.mod .
COPY main.go .
RUN go mod download

COPY . .

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]