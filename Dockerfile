FROM golang:1.13 as build

WORKDIR /sc
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /sc/status -a -installsuffix cgo main.go

#FROM scratch
FROM golang:1.13

COPY --from=build /sc/status /sc/status

CMD [ "/sc/status" ]
