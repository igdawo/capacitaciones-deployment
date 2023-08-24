FROM acurilem/backend-veterinaria:latest

WORKDIR /app

RUN go build -o main ./main.go

EXPOSE 3001
CMD [ "/app/main" ]