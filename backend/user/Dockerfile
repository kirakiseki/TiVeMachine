# build stage
FROM --platform=$BUILDPLATFORM golang:1.20.4 as build

ARG TARGETARCH
ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN go mod tidy

RUN if [ "$TARGETARCH" == "arm64" ]; then CC=aarch64-linux-gnu-gcc && CC_FOR_TARGET=gcc-aarch64-linux-gnu && apt-get update && apt-get install -y gcc-aarch64-linux-gnu; fi && \
  CGO_ENABLED=1 GOOS=linux GOARCH=$TARGETARCH CC=$CC CC_FOR_TARGET=$CC_FOR_TARGET go build -a -ldflags '-extldflags "-static"' -o /main main.go

# production stage
FROM alpine as production

ENV TZ=Asia/Shanghai
LABEL maintainer="Ishirai <ishirai@163.com>"

WORKDIR /app
COPY --from=build /app/main .

ENV PORT=80\
    DSN=root:root@tcp(traefik:3306)/TiVeMachine?charset=utf8mb4&parseTime=True&loc=Local

EXPOSE 80

ENTRYPOINT [ "./main" ]