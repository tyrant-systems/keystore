FROM golang:alpine as build
WORKDIR /artifacts
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o keystore main.go

FROM alpine:latest
ARG port=3030
ARG config=config.json
WORKDIR /srv
COPY --from=build /artifacts/keystore keystore
COPY ${config} cfg.json
EXPOSE ${port}
ENTRYPOINT [ "./keystore", "-config", "cfg.json" ]
