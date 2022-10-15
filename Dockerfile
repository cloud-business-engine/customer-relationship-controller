FROM golang:1.18.3-bullseye as builder

ARG GIT_HOST=git.ucb.local

WORKDIR /go/deleivery
COPY . .

RUN go build -o ./ ./cmd/http_api

FROM golang:1.18.3-bullseye as runtime

ENV TZ=Europe/Moscow
ENV GIN_MODE=release

WORKDIR /app/http_api

COPY --from=builder /go/deleivery/http_api ./

RUN useradd -r -UM app
USER app

CMD ["./http_api"]
