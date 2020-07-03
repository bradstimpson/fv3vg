FROM golang:alpine as builder
RUN mkdir /build 
ADD ./gin_server /build/
WORKDIR /build
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch 
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 8002
CMD ["./main"]