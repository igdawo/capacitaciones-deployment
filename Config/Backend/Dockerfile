FROM golang:1.21.1-alpine

WORKDIR /app

ARG BRANCH_BACK
ARG GITHUB_TOKEN
RUN git clone -b ${BRANCH_BACK} https://${GITHUB_TOKEN}@github.com/citiaps/capacitaciones-back-go.git /app

RUN go build -o main ./main.go

CMD [ "/app/main" ]