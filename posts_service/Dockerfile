FROM golang:1.21.1 AS build

COPY . /go/src/posts_service
WORKDIR /go/src/posts_service

ENV CGO_ENABLED=0
RUN go mod download -x

RUN go build -o swagger .

FROM scratch AS runtime
COPY --from=build /go/src/posts_service/swagger ./swagger
ENTRYPOINT ["./swagger"]
