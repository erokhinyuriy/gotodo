FROM golang:alpine AS builder

WORKDIR /usr/local/src

#dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

#build
COPY . .
RUN go build -o main main.go


FROM alpine

WORKDIR /usr/local/src

COPY --from=builder /usr/local/src/ /usr/local/src/

EXPOSE 8447

RUN chmod a+x .

CMD ["./main"]