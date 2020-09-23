From golang:1.15.2

WORKDIR ~/go/src/cns-auth

COPY . .

RUN go get ./...

RUN go install ./...

CMD ["cns-auth"]


