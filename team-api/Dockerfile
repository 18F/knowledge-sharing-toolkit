FROM dev-standard

ENV _TEAM_API_SERVER_VERSION 0.1.1
RUN bash -l -c "npm install -g team-api-server@$_TEAM_API_SERVER_VERSION"

RUN groupadd -r team-api && \
    useradd -r -m -g team-api team-api && \
    mkdir $APP_SYS_ROOT/team-api && \
    chown team-api:team-api $APP_SYS_ROOT/team-api
USER team-api-server
WORKDIR $APP_SYS_ROOT/team-api
VOLUME $APP_SYS_ROOT/team-api
COPY ["team-api-config.json", "team-api-config.json"]
EXPOSE 6000 6001
CMD ["/bin/bash", "-l", "-c", "team-api-server team-api-config.json"]
