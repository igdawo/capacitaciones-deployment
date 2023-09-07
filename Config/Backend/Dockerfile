FROM golang:1.21.0

WORKDIR /app

ARG GITHUB_TOKEN
ARG BRANCH

RUN git clone -b ${BRANCH} https://${GITHUB_TOKEN}@github.com/citiaps/capacitaciones-back-go.git /app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change

RUN go mod download 
RUN go get github.com/gin-gonic/gin/binding@v1.8.2

RUN go build -o main .

EXPOSE 3001
CMD ["./main"]