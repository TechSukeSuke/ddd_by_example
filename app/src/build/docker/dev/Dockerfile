# ==================================================
# アプリケーションで必要なものを最小限に留めたレイヤー
# ==================================================
FROM golang:1.21.0-alpine AS deps

RUN apk add --update --no-cache bash git make

ENV GOBIN=/go/bin

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2 && \
    go install github.com/go-delve/delve/cmd/dlv@v1.22.1 && \
    go install github.com/google/wire/cmd/wire@v0.6.0 && \
    go install github.com/cosmtrek/air@v1.52.0

# ==================================================
# アプリケーションの実行レイヤー
# ==================================================
FROM golang:1.21.0-alpine AS runner

RUN apk add --update --no-cache bash git make

ENV GO111MODULE=on
ENV GOBIN=/go/bin

WORKDIR /go/bin

COPY --from=deps /go/bin .

# ※ go mod tidyを実行する前にソースコードをコピーすること。※
# import部分を見てgo.modを更新するため、必要最低限のパッケージしか取得できない。
WORKDIR /go/src

COPY . .

CMD ["air", "-c", "/go/src/.air.toml"]
