FROM postgres:latest

ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=veterinaria_citiaps

COPY veterinaria_citiaps.zip /

RUN apt-get update && apt-get install -y unzip

RUN unzip veterinaria_citiaps.zip

#WORKDIR /backup

EXPOSE 5432

COPY restore-database.sh /docker-entrypoint-initdb.d/restore-database.sh

RUN chmod +x /docker-entrypoint-initdb.d/restore-database.sh




