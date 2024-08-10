FROM golang:1.22.2 as builder

WORKDIR /app

COPY . /app/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server /app/cmd/talison_contabil/main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /server /server
COPY --from=builder /app/config.toml /config.toml

EXPOSE 3033

CMD [ "/server" ]
