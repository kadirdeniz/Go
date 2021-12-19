FROM golang:1.17-buster as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -o /server
EXPOSE 3000
CMD ["/server"]