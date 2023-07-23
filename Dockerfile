FROM golang:latest AS builder

WORKDIR /emitter
COPY . .

RUN apt install git \
    && go build -o emitter . \
    && chmod +x emitter


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /emitter

COPY --from=builder /emitter/emitter .

# Main port
EXPOSE 8080
# Cluster port
EXPOSE 4000


CMD ["./emitter"]
