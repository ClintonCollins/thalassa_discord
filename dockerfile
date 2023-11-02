# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20.5-alpine AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o thalassa_discord ./cmd/thalassa_discord


FROM alpine AS build-release

WORKDIR /

RUN apk update --no-cache
RUN apk upgrade --no-cache
RUN apk --no-cache add ffmpeg
RUN apk --no-cache add yt-dlp

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/thalassa_discord /app/thalassa_discord

EXPOSE 6500

ENTRYPOINT ["/app/thalassa_discord"]