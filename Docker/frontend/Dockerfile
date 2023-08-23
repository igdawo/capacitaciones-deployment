# Establece la imagen base
FROM node:14


RUN npm install -g http-server
# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de la aplicación al directorio de trabajo
COPY . .

# Instala las dependencias
RUN npm install

# Expone el puerto 3000 (puerto por defecto de Nuxt)
EXPOSE 3000

RUN npm run build

# Comando para ejecutar el servidor de desarrollo de Nuxt
CMD ["npm", "run","dev"]