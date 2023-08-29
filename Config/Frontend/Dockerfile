FROM node:16

# Establecer variables de entorno
ARG GITHUB_TOKEN
ARG BRANCH=fmiranda

# Clonar el repositorio y la rama específica
RUN git clone -b ${BRANCH} https://${GITHUB_TOKEN}@github.com/citiaps/capacitaciones-front-nuxt.git /app

# Cambiar al directorio de trabajo
WORKDIR /app

# Instalar dependencias
#COPY package*.json ./
RUN npm install

# Copiar otros archivos y construir
#COPY . .
RUN npm run build

# Exponer puerto
EXPOSE 3000

# Comando para iniciar la aplicación
CMD ["npm", "run", "dev"]