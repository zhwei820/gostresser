
FROM xqdocker/ubuntu-nginx
MAINTAINER HelloFresh
WORKDIR /code
COPY ./ui/deploy/default.conf /etc/nginx/conf.d/default.conf
COPY ./ui/dist/ /usr/share/nginx/html/
COPY ./ui/public/externalConfig.js.tmpl /tmp/
EXPOSE 80
CMD envsubst < /tmp/externalConfig.js.tmpl > /usr/share/nginx/html/externalConfig.js && nginx -g 'daemon off;'

COPY ./ui/deploy/start.sh ./
COPY ./gostresser ./
COPY ./worker ./worker/
COPY ./config ./config/

# docker build . -t daocloud.io/zhwei820/gostresser
