# FROM golang:1.17
# RUN mkdir -p /app 
# WORKDIR /app
# ADD . /app
# RUN go mod download
# RUN go build ./main.go
# CMD ["./main"]

FROM golang:1.18-alpine as builder
WORKDIR /go/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app main.go
FROM gcr.io/distroless/base
COPY --from=builder /go/app/ .
CMD ["/app"]