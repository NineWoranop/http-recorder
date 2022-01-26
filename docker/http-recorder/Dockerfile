FROM debian:bullseye-20211011
LABEL Vendor="http-recorder" \
      maintainer="NineWoranop@users.noreply.github.com"

ADD https://github.com/NineWoranop/http-recorder/releases/download/v0.1/http-recorder-linux-amd64.tar.gz  /

RUN cd /;tar -xvzf http-recorder-linux-amd64.tar.gz;rm -rf http-recorder-linux-amd64.tar.gz;

# start service
ENTRYPOINT ["/http-recorder-linux-amd64"]
