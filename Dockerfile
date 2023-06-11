# ビルドステージ1: Goコードのビルド用のベースイメージ（例: golang:1.16-buster）
FROM golang:1.19-buster AS build-stage

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp

# ビルドステージ2: 実行環境用のベースイメージ（例: alpine:3.14）
FROM alpine:3.14 AS runtime-stage

WORKDIR /app

COPY --from=build-stage /app/myapp .

CMD ["./myapp"]
