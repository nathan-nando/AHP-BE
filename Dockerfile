FROM golang:1.18-alpine AS build
RUN apk update && \
    apk add curl \
    git \
    bash \
    make \
    ca-certificates && \
    rm -rf /var/cache/apk/*


WORKDIR /app

ENV GOPATH /go
COPY go.* ./
RUN go mod download
RUN go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o application main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash tzdata
ENV TZ Asia/Jakarta

WORKDIR /app/
COPY --from=build /app/application .
COPY --from=build /app/assets ./assets/
COPY --from=build /app/config/prod.yaml ./config/

RUN ["chmod", "+x", "./application"]
ENTRYPOINT ["./application"]

