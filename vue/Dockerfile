FROM node:12.12.0-alpine

WORKDIR /frontend

COPY ./frontend/package.json .
COPY ./frontend/package-lock.json .

RUN apk update && \
    npm install -g npm @vue/cli \
    npm install