FROM mongo:5.0

COPY veterinaria_citiaps.zip /

RUN apt-get update && apt-get install -y unzip

RUN unzip veterinaria_citiaps.zip

WORKDIR /veterinaria_citiaps

EXPOSE 27017

COPY restore-database.sh /docker-entrypoint-initdb.d/restore-database.sh

RUN chmod +x /docker-entrypoint-initdb.d/restore-database.sh