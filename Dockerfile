FROM node:10.15.1-jessie as builder
WORKDIR /code
COPY . ./
WORKDIR /code/ui

RUN yarn install
RUN yarn build

# -----

FROM xqdocker/ubuntu-nginx
MAINTAINER HelloFresh
WORKDIR /code
COPY --from=builder /code/ui/deploy/default.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /code/ui/dist/ /usr/share/nginx/html/
COPY --from=builder /code/ui/public/externalConfig.js.tmpl /tmp/
EXPOSE 80
CMD envsubst < /tmp/externalConfig.js.tmpl > /usr/share/nginx/html/externalConfig.js && nginx -g 'daemon off;'

COPY --from=builder /code/ui/deploy/start.sh ./
COPY --from=builder /code/gostresser ./
COPY --from=builder /code/worker ./worker/
COPY --from=builder /code/config ./config/

# docker build . -t daocloud.io/zhwei820/gostresser
