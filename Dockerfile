FROM golang:1.14 as builder
LABEL maintainer="Krisna Putra <anakgembala21@gmail.com>"
ENV GO111MODULE=on
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy everything from the current directory to the Working Directory inside the container
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/ .

EXPOSE 8082
CMD ["./main"]