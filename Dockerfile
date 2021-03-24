FROM golang:1.14-alpine AS build

WORKDIR /src/
COPY . /src/
RUN apk update
RUN apk add git
RUN go get -v -d ./...
RUN CGO_ENABLED=0 go build -o /bin/cape-api

FROM alpine 
WORKDIR /project
COPY --from=build /bin/cape-api /project/cape-api
COPY static /project/static
ENTRYPOINT ["/project/cape-api"]
