FROM golang:1.20.5-alpine

WORKDIR /go/src/app

COPY ./food_delivery_api ./food_delivery_api



FROM node:18-alpine

WORKDIR /srv/app

COPY ./food_delivery-frontend-v2 /srv/app

COPY entrypoint.sh /srv/app/entrypoint.sh

RUN chmod +x /srv/app/entrypoint.sh

EXPOSE 8000

CMD ["/srv/app/entrypoint.sh"]



