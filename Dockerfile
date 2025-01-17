FROM golang:1.18.1

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN env GOOS=linux GOARCH=amd64 go build -o thourus-api

EXPOSE 50051

CMD [ "./thourus-api" ]