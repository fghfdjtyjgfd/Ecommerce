FROM golang:1.22.2-bullseye as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:3.18 as final
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 8000
CMD [ "./main" ]
