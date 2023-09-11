FROM node:18.16.0
WORKDIR /usr/src/app

ARG GITHUB_TOKEN
ARG BRANCH
RUN git clone -b ${BRANCH} https://${GITHUB_TOKEN}@github.com/citiaps/capacitaciones-front-nuxt.git /usr/src/app
#RUN git clone -b ${BRANCH} https://github.com/citiaps/capacitaciones-front-nuxt.git /usr/src/app
RUN mv .env.example .env
#COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3000
# Comando para ejecutar el servidor Nuxt.js cuando se inicia el contenedor
#CMD [ "npm", "start" ]
CMD [ "npm", "run", "dev" ]
