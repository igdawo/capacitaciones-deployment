FROM node:20.5.1

WORKDIR /app

ARG GITHUB_TOKEN
ARG BRANCH

RUN git clone -b ${BRANCH} https://${GITHUB_TOKEN}@github.com/citiaps/capacitaciones-front-nuxt.git /app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change

#RUN npm run build
RUN npm install

#RUN npm run dev


EXPOSE 3000
CMD [ "npm", "run","dev" ]