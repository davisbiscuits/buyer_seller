FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY wait-for.sh .

FROM alpine:3.16
WORKDIR /app

# Copy the built binary and scripts from the builder stage
COPY --from=builder /app/main .
COPY wait-for.sh .

# Make wait-for.sh executable
RUN chmod +x wait-for.sh

EXPOSE 8080
CMD [ "./wait-for.sh", "db:3306", "--", "./main" ]