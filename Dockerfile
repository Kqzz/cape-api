FROM alpine:3.7
ADD cape-api server
ENV PORT 80
EXPOSE 80
ENTRYPOINT ["/server"]