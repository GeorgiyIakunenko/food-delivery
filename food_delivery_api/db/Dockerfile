FROM postgres:14.8-alpine

ENV POSTGRES_DB=fooddelivery
ENV POSTGRES_USER=fooddelivery
ENV POSTGRES_PASSWORD=123456

COPY . /docker-entrypoint-initdb.d/

EXPOSE 5432