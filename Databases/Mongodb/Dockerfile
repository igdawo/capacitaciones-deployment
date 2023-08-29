FROM mongo:latest
ADD backup_veterinaria_citiaps.tar.gz /backup/
COPY restore.sh /docker-entrypoint-initdb.d/
RUN chmod +x /docker-entrypoint-initdb.d/restore.sh