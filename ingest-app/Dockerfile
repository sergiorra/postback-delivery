FROM php:7.4-apache

LABEL maintainer="sergiorra17@gmail.com"
LABEL application="ingest-app"

RUN apt-get update && \
    pecl install redis-5.3.2 && \
    docker-php-ext-enable redis

ENV DISPLAY_ERRORS="On"

COPY ingest-app/php.ini /usr/local/tc/php/
COPY ingest-app/index.php /var/www/html/