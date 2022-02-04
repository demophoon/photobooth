# Go Shrls
FROM golang:1.16-alpine AS builder
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY ./server/*.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a --installsuffix cgo -o photobooth

# Shrls Frontend
FROM node:17-alpine as frontend
WORKDIR /app

COPY . /app
RUN npm install
RUN npm run build-prod

# Final Artifact
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/photobooth ./
COPY --from=frontend /app/dist/ /dist/

CMD ["/photobooth"]
