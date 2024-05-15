FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o api .

EXPOSE 4242

ENTRYPOINT [ "/app/api" ]