FROM golang:1.23.1 as build

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /go/bin/app

FROM scratch

COPY --from=build /go/bin/app /

CMD [ "/app" ]
