FROM dev-standard

RUN bash -l -c 'go get github.com/bitly/oauth2_proxy && \
    strip $GOPATH/bin/oauth2_proxy'

USER root
RUN groupadd -r oauth2_proxy && useradd -r -m -g oauth2_proxy oauth2_proxy
USER oauth2_proxy
WORKDIR $APP_SYS_ROOT/oauth2_proxy

EXPOSE 4180
COPY ["entrypoint.sh", "entrypoint.sh"]
ENTRYPOINT ["./entrypoint.sh"]
CMD ["run-server"]
