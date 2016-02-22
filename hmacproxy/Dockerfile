FROM dev-standard

RUN bash -l -c 'go get github.com/18F/hmacproxy && \
    strip $GOPATH/bin/hmacproxy'

RUN groupadd -r hmacproxy && useradd -r -m -g hmacproxy hmacproxy && \
    mkdir $APP_SYS_ROOT/hmacproxy && \
    chown hmacproxy:hmacproxy $APP_SYS_ROOT/hmacproxy
USER hmacproxy
WORKDIR $APP_SYS_ROOT/hmacproxy
EXPOSE 8083
CMD ["/bin/bash", "-l", "-c", \
     "$GOPATH/bin/hmacproxy -auth -port 8083 -secret $HMACAUTH_SECRET \
      -sign-header Team-Api-Signature -headers Content-Type,Date"]
