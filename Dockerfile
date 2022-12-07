# 1. build
FROM golang:1.19 AS build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# 2. run
FROM alpine
WORKDIR /app
COPY --from=build ./build/server .
EXPOSE 1323
CMD ["./server"]
