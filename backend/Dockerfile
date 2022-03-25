FROM golang:1.18
RUN mkdir /app

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV PORT 3000
RUN go mod tidy
RUN go build -o server .
CMD ["/app/server"]
