FROM golang:1.21.6 as build

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/base-debian12

COPY --from=build /go/bin/app /
COPY --from=build /app/content /content
COPY --from=build /app/pages /pages
COPY --from=build /app/public /public

CMD [ "/app" ]
