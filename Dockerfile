FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /pendown-be

EXPOSE 8080

CMD [ "/pendown-be" ]