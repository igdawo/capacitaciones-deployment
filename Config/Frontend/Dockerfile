FROM node:lts

WORKDIR /app

ARG BRANCH_FRONT
ARG GITHUB_TOKEN
RUN git clone -b ${BRANCH_FRONT} https://${GITHUB_TOKEN}@github.com/citiaps/capacitaciones-front-nuxt.git /app

RUN npm install

RUN apt update
RUN apt install vim -y
RUN apt install iputils-ping -y

ENV HOST 0.0.0.0
CMD [ "yarn", "dev" ]