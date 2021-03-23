FROM golang:1.14-alpine AS build

WORKDIR /src/
COPY *.go /src/
RUN apk update
RUN apk add git
RUN go get -v -d ./...
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM alpine 
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]
