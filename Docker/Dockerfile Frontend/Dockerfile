# Usar una imagen base de Node.js
FROM node:16.14.2

# Establecer el directorio de trabajo en el contenedor
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

# Copiar el package.json y package-lock.json al directorio de trabajo
COPY package*.json ./

# Instalar las dependencias
RUN npm install

# Copiar el resto de los archivos de la aplicación
COPY . .

# Construir la aplicación Nuxt
RUN npm run build

# Exponer el puerto en el que Nuxt está escuchando (por defecto: 3000)
EXPOSE 3000

# Comando para ejecutar la aplicación Nuxt
CMD [ "npm", "start"]