FROM golang:alpine as builder

COPY cmd /app/cmd
COPY internal /app/internal
COPY go.mod /app

WORKDIR /app

RUN go mod tidy && \
    go build -o willdo ./cmd/willdo && \
    go clean -cache -modcache

FROM alpine
RUN apk add --no-cache ca-certificates && \
    rm -rf /var/cache/apk/*
COPY --from=builder /app/willdo willdo

RUN addgroup -S app_user_group && \
    adduser -S app_user -G app_user_group && \
    chown -R app_user:app_user_group willdo
USER app_user

EXPOSE 8080
ENTRYPOINT ["./willdo"]