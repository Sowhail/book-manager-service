FROM golang:1.20.7-alpine

ENV GO111MODULE=on
ENV GOOS=linux
ENV CGO_ENABLED=0

WORKDIR /bookManager

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

EXPOSE 3001

CMD [ "go", "run", "." ]