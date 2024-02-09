FROM golang:1.22.0 as build

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /go/bin/app

FROM gcr.io/distroless/base-debian12

COPY --from=build /go/bin/app /

CMD [ "/app" ]
