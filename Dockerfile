FROM golang:1.21-alpine AS build

WORKDIR /src/
COPY . /src/
RUN apk --no-cache update
# RUN apk --no-cache add git
RUN go get -v -d ./...
RUN CGO_ENABLED=0 go build -o /bin/cape-api

FROM alpine 
WORKDIR /project
COPY --from=build /bin/cape-api /project/cape-api
COPY static /project/static
ENTRYPOINT ["/project/cape-api"]
