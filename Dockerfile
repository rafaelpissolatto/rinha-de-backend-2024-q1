FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o app cmd/api/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/app /
COPY --from=builder /app/.env /
CMD ["/app"]