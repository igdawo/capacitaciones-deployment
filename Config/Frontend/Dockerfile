FROM node:18.16.0
WORKDIR /usr/src/app
ARG GITHUB_TOKEN
ARG BRANCH
RUN git clone -b ${BRANCH} https://${GITHUB_TOKEN}@github.com/citiaps/capacitaciones-front-nuxt.git /usr/src/app
RUN npm install
EXPOSE 3000
CMD [ "npm", "run", "dev" ]

#CMD [ "npm", "run", "build" ]
#CMD [ "npm", "start" ]
