FROM dev-standard

RUN bash -l -c 'go get github.com/bitly/oauth2_proxy && \
    strip $GOPATH/bin/oauth2_proxy'

RUN groupadd -r oauth2_proxy && useradd -r -m -g oauth2_proxy oauth2_proxy && \
    mkdir $APP_SYS_ROOT/oauth2_proxy && \
    chown oauth2_proxy:oauth2_proxy $APP_SYS_ROOT/oauth2_proxy
USER oauth2_proxy
WORKDIR $APP_SYS_ROOT/oauth2_proxy

COPY ["oauth2_proxy.cfg", "oauth2_proxy.cfg"]
EXPOSE 4180
CMD ["/bin/bash", "-l", "-c", \
     "$GOPATH/bin/oauth2_proxy --config=oauth2_proxy.cfg"]
