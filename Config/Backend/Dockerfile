FROM golang:1.21.0

WORKDIR /app

# Establecer variables de entorno
ARG GITHUB_TOKEN
ARG BRANCH

# Clonar el repositorio
RUN git clone -b ${BRANCH} https://${GITHUB_TOKEN}@github.com/citiaps/capacitaciones-back-go.git /app

# Instalar dependencias
RUN go mod download
RUN go get github.com/gin-gonic/gin/binding@v1.8.2

# Construir la aplicación
RUN go build -o main .

# Exponer el puerto
EXPOSE 3001

# Comando para iniciar la aplicación
CMD ["./main"]
