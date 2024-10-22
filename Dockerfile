FROM golang:1.23.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /go/bin/app

FROM scratch

COPY --from=build /go/bin/app /app

ENTRYPOINT ["/app"]
