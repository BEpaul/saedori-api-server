FROM golang:1.24.2-alpine

WORKDIR /app
COPY . .

RUN go mod download

WORKDIR /app/cmd

CMD ["go", "run", "."]
