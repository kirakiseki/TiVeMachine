# build stage
FROM --platform=$BUILDPLATFORM tonistiigi/xx AS xx

FROM --platform=$BUILDPLATFORM golang:1.20.4-alpine as build

COPY --from=xx / /

WORKDIR /app
COPY . .

RUN apk add clang lld
ARG TARGETPLATFORM
RUN xx-apk add musl-dev gcc
ENV CGO_ENABLED=1
ENV GO111MODULE=on

RUN go mod tidy
RUN xx-go build -o main ./main.go && \
    xx-verify main

# production stage
FROM alpine as production

ENV TZ=Asia/Shanghai
LABEL maintainer="Ishirai <ishirai@163.com>"

WORKDIR /app
COPY --from=build /app/main ./main

ENV PORT=80\
    DSN=root:root@tcp(traefik:3306)/TiVeMachine?charset=utf8mb4&parseTime=True&loc=Local

EXPOSE 80

ENTRYPOINT [ "./main" ]