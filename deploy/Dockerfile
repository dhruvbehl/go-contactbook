FROM golang:1.16-alpine

WORKDIR /app

COPY cmd/*.go ./

RUN go build -o /contact-book

EXPOSE 8080

CMD [ "/contact-book" ]