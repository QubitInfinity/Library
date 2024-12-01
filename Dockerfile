# Deploy the application binary into a lean image
FROM golang:1.23

WORKDIR /

COPY .  .

RUN go build -o main

EXPOSE 8080

ENTRYPOINT ["/main"]