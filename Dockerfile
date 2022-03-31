FROM golang:alpine
ENV CGO_ENABLED=0

RUN apk add git

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main .

EXPOSE $PORT

CMD [ "./main" ]