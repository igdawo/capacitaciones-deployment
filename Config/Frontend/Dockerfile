FROM node:16.10.0-alpine

WORKDIR /app

RUN apk update && apk upgrade
RUN apk add git

ARG BRANCH_FRONT
ARG GITHUB_TOKEN
RUN git clone -b ${BRANCH_FRONT} https://${GITHUB_TOKEN}@github.com/citiaps/capacitaciones-front-nuxt.git /app
COPY .env.docker .env

RUN npm install
RUN npm run build

RUN apk add vim
RUN apk add iputils

ENV NUXT_HOST=0.0.0.0
ENV NUXT_PORT=3000
CMD [ "yarn", "start" ]
