FROM golang:1.22.2-alpine

WORKDIR /httpapp

COPY . .

RUN go build /httpapp

EXPOSE 8000

ENTRYPOINT ["./videos"]

