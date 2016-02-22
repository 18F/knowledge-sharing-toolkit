FROM dev-standard

ENV _LUNR_SERVER_VERSION 0.1.3
RUN bash -l -c "npm install -g lunr-server@$_LUNR_SERVER_VERSION"

RUN groupadd -r lunr-server && useradd -r -m -g lunr-server lunr-server && \
    mkdir $APP_SYS_ROOT/lunr-server && \
    chown lunr-server:lunr-server $APP_SYS_ROOT/lunr-server
USER lunr-server
WORKDIR $APP_SYS_ROOT/lunr-server
VOLUME $APP_SYS_ROOT/pages/sites
VOLUME $APP_SYS_ROOT/hub/repo
COPY ["./lunr-server-config.json", "lunr-server-config.json"]
EXPOSE 7777
CMD ["/bin/bash", "-l", "-c", "lunr-server lunr-server-config.json"]
