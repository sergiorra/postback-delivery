FROM php:7.4-apache

LABEL maintainer="sergiorra17@gmail.com"
LABEL application="ingest-app"

RUN apt-get update && \
    pecl install redis-5.3.2 && \
    docker-php-ext-enable redis

COPY ingest-app/index.php /var/www/html/