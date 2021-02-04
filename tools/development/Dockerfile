FROM golang:1.15.7-alpine3.13

RUN apk --update add --no-cache make nginx

ENV APP_DIR=/opt/service
WORKDIR $APP_DIR

COPY . $APP_DIR
COPY ./tools/development/nginx/nginx.conf /etc/nginx/conf.d/default.conf

RUN mkdir -p /run/nginx

ENTRYPOINT ["/opt/service/tools/development/run"]

