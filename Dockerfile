FROM golang:1.21-alpine as builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./transactionapi ./main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/transactionapi ./transactionapi
COPY config.json .
CMD [ "/app/transactionapi" ]