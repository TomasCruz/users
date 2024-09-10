FROM golang:1.23
WORKDIR /app
COPY . .
COPY .env .
RUN go mod download
RUN go build -o /app/bin/server cmd/*
EXPOSE 8091
CMD ["/app/bin/server"]