FROM golang:1.8.3-alpine
LABEL maintainer "sambalana247@gmail.com"

# Download some needed tools
RUN apk add --no-cache --virtual bootstrap-deps git wget ca-certificates \
    # Download the wait script
    && wget https://github.com/Eficode/wait-for/raw/master/wait-for -P /bin \
    && chmod +x /bin/wait-for \
    # Download the hot reload tool
    && go get github.com/tockins/realize \
    # Clean up
    && apk del bootstrap-deps

COPY . /go/src/github.com/puradox/schemed
WORKDIR /go/src/github.com/puradox/schemed

CMD realize run
