FROM golang:alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .

EXPOSE 8000

CMD ["air"]