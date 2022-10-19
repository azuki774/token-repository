FROM alpine:3.7
ADD build/bin/token-repository /usr/local/bin/token-repository
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
ENTRYPOINT ["/usr/local/bin/token-repository", "start"]
EXPOSE 80
