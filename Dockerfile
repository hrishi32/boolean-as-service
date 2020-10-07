# FROM golang:latest as builder
# WORKDIR /app
# COPY . .
# RUN go mod download
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# EXPOSE 8000
# CMD ["./main"]

# FROM mysql:latest

FROM golang:alpine as builder

LABEL maintainer="Hrushikesh Sarode <sarodehrishikesh18@gmail.com>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

# RUN go build .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8000


CMD ["./main"]