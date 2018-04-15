FROM    alpine:3.3
ENV     PORT 3000

WORKDIR /var/www/chat/

RUN apk --no-cache add ca-certificates

EXPOSE  ${PORT}
CMD     exec bin/server

ADD     bin/server bin/server
