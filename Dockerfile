FROM golang:latest AS builder

WORKDIR /app
COPY . .

RUN apt install git \
    && go build -o emitter . \
    && chmod +x emitter


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/emitter .

RUN chmod +x /app/emitter

# Main port
EXPOSE 8080
# Cluster port
EXPOSE 4000


CMD ["/app/emitter"]
