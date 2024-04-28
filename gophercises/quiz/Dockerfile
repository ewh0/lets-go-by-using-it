FROM golang:alpine

RUN mkdir -p /root/quiz/

WORKDIR /root/quiz/

COPY quiz.go /root/quiz/
COPY problems.csv /root/quiz/
COPY go.mod /root/quiz/

RUN go build .

ENTRYPOINT [ "./quiz" ]