FROM node:lts as node_builder
WORKDIR /site
COPY tailwind.config.js ./
COPY web/templates ./web/templates/
COPY web/styles ./web/styles/
COPY package*.json ./
RUN npm ci
RUN npm run build-styles

FROM golang:1.18 as go_builder
WORKDIR /site
COPY go.* ./
RUN go mod download
COPY cmd ./cmd/
COPY internal ./internal/
RUN CGO_ENABLED=0 go build -o build/site cmd/site/main.go

FROM alpine:latest
WORKDIR /site
COPY site.dhall ./
COPY web/templates ./web/templates/
COPY web/static/img ./build/static/img/
COPY blog ./blog/
COPY --from=go_builder /site/build/site ./build/
COPY --from=node_builder /site/build/app.css ./build/static/css/
CMD ["./build/site"]

