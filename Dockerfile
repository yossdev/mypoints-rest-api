# Initial stage: download modules and build environment
FROM golang:1.17-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o app

# Stage 2:
FROM alpine:3.14
WORKDIR /root/
COPY --from=builder /app/.env .
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]