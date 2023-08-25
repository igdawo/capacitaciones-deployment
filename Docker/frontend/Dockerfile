# Dockerfile
FROM node:alpine

EXPOSE 3000 8080

# create destination directory
RUN mkdir -p /usr/src/nuxt-app
WORKDIR /usr/src/nuxt-app

# update and install dependency
RUN apk update && apk upgrade
#RUN apk add git

# copy the app, note .dockerignore
COPY . /usr/src/nuxt-app/
RUN npm install
RUN npm run build

CMD [ "npm", "start" ]

#Para construir contenedor correr:
# docker build -t fondoleon_front .
# docker run -d --name admin-leon-frontend --network admin-leon_default -p 3000:3000 fondoleon_front
