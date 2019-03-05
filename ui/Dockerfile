FROM node:10.15.1-jessie as builder
WORKDIR /code

COPY . ./
RUN yarn install
RUN yarn build

# -----

FROM nginx:stable-alpine
MAINTAINER HelloFresh
COPY --from=builder /code/deploy/default.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /code/dist/ /usr/share/nginx/html/
COPY --from=builder /code/public/externalConfig.js.tmpl /tmp/
EXPOSE 80
CMD envsubst < /tmp/externalConfig.js.tmpl > /usr/share/nginx/html/externalConfig.js && nginx -g 'daemon off;'

# docker build . -t daocloud.io/zhwei820/GoStresser
