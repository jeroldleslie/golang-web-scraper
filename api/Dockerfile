FROM golang AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY ./api/go.mod .
COPY ./api/go.sum .

RUN go mod download

COPY . .

RUN cd /app/api && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

# final stage
FROM scratch AS final
COPY --from=builder /app/api /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose both 443 and 80 to our application
EXPOSE 443
EXPOSE 80

# Mount the certificate cache directory as a volume, so it remains even after
# we deploy a new version
VOLUME ["/cert-cache"]

CMD ["./app/server"]