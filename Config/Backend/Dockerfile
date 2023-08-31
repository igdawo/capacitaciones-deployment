FROM golang:1.21.0
WORKDIR /app

ARG GITHUB_TOKEN
ARG BRANCH
RUN git clone -b ${BRANCH} https://${GITHUB_TOKEN}@github.com/citiaps/capacitaciones-back-go.git /app
#COPY .env .env
COPY . .
RUN go mod download
EXPOSE 3000
CMD ["go", "run", "main.go"]