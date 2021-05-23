FROM golang:alpine as builder

ENV GO111MODULE=on

WORKDIR /app 

COPY . .

COPY go.mod .

COPY go.sum .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags="-w -s" .

FROM scratch

WORKDIR /app

COPY --from=builder /app/api /usr/bin/

ENTRYPOINT ["api"]
