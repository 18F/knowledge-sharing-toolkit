FROM dev-standard

ENV _LUNR_SERVER_VERSION 0.1.3
RUN bash -l -c "npm install -g lunr-server@$_LUNR_SERVER_VERSION"

RUN groupadd -r lunr-server && useradd -r -m -g lunr-server lunr-server
USER lunr-server
WORKDIR $APP_SYS_ROOT/lunr-server

EXPOSE 7777
CMD ["/bin/bash", "-l", "-c", "lunr-server config/lunr-server-config.json"]
