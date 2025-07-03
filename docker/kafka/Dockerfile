FROM golang:1.24
WORKDIR /app
COPY . .
RUN go build -o /init ./cmd/init
CMD ["/init"]