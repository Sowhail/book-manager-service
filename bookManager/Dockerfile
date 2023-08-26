# First Stage
FROM golang:1.20.7-alpine as firstStage

ENV GO111MODULE=on
ENV GOOS=linux
ENV CGO_ENABLED=0

WORKDIR /bookManager

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

EXPOSE 3001

RUN go build -o result .


# second Stage Execution
FROM alpine:3.18.2

WORKDIR /bookManager

COPY --from=firstStage /bookManager/result .

CMD [ "./result" ]
