# build docker image
FROM alpine:latest
COPY go-starwars /usr/local/bin/go-starwars

ENTRYPOINT ["go-starwars"]
