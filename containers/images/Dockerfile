FROM golang:1.20-alpine
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /bin/app ./cmd

FROM alpine
WORKDIR /usr/src/app
COPY --from=0 /bin/app /bin/app
COPY --from=0 src/views /usr/src/app/views
COPY --from=0 src/assets /usr/src/app/assets


ENTRYPOINT [ "/bin/app" ]
