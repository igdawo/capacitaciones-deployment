# Utiliza una imagen base de Node.js
FROM node:14-alpine

# Establece el directorio de trabajo en la aplicación
WORKDIR /app

COPY . .

RUN npm install

EXPOSE 3000

ENV NUXT_HOST=0.0.0.0

ENV NUXT_PORT=3000

CMD [ "npm", "run", "dev"]
