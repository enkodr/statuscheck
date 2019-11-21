FROM golang:1.13 as build

WORKDIR /hc
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o /hc/healthcheck -a -installsuffix cgo main.go

#FROM scratch
FROM golang:1.13

COPY --from=build /hc/healthcheck /hc/healthcheck

CMD [ "/hc/healthcheck" ]
