FROM golang:1.15-alpine AS GO_BUILD
COPY . /bookstore_users-api
WORKDIR /bookstore_users-api/bookstore_users-api
RUN go build -o /go/bin/bookstore_users-api/bookstore_users-api

FROM alpine:3.10
WORKDIR app
COPY --from=GO_BUILD /go/bin/bookstore_users-api/ ./
EXPOSE 8080
CMD ./bookstore_users-api