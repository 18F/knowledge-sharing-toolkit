FROM dev-standard

RUN bash -l -c 'go get github.com/18F/hmacproxy && \
    strip $GOPATH/bin/hmacproxy'

USER root
RUN groupadd -r hmacproxy && useradd -r -m -g hmacproxy hmacproxy
USER hmacproxy
WORKDIR $APP_SYS_ROOT/hmacproxy

EXPOSE 8083 8084
COPY ["entrypoint.sh", "entrypoint.sh"]
ENTRYPOINT ["./entrypoint.sh"]
CMD ["run-server"]
