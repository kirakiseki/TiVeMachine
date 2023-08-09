# build stage
FROM node:16.20.1-alpine3.18 as build

WORKDIR /app
COPY ../../frontend/package*.json .
RUN npm install
COPY ../../frontend .
RUN npm run build

# production stage
FROM nginx:stable-alpine3.17-slim as production

ENV TZ=Asia/Shanghai
LABEL maintainer="Ishirai <ishirai@163.com>"
COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80