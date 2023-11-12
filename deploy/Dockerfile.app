FROM golang:alpine

WORKDIR /build

COPY ./resource/app.yml ./resource/app.yml
COPY go.mod go.sum ./
COPY ./cmd/app ./cmd/app/
COPY ./internal ./internal/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/app

EXPOSE 8080

CMD [ "/app" ]