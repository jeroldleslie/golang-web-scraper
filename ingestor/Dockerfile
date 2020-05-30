FROM golang AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY ./ingestor/go.mod .
COPY ./ingestor/go.sum .

RUN go mod download

COPY . .

RUN cd /app/ingestor && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

# final stage
FROM scratch AS final
COPY --from=builder /app/ingestor /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose both 443 and 80 to our application
EXPOSE 443
EXPOSE 80

# Mount the certificate cache directory as a volume, so it remains even after
# we deploy a new version
VOLUME ["/cert-cache"]

CMD ["./app/main"]