FROM alpine:latest

WORKDIR /app

COPY build/sus-gopher-facts /app

EXPOSE 8080

CMD ["./sus-gopher-facts"]