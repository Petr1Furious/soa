FROM golang:1.21.1 AS build

COPY main_service /go/src/main_service
COPY posts_service /go/src/posts_service
COPY go* /go/src/
WORKDIR /go/src/main_service

ENV CGO_ENABLED=0
RUN go mod download -x

RUN go build -o swagger .

FROM scratch AS runtime
COPY --from=build /go/src/main_service/swagger ./swagger
ENTRYPOINT ["./swagger"]
