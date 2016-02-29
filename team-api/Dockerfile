FROM dev-standard

ENV _TEAM_API_SERVER_VERSION 0.1.1
RUN bash -l -c "npm install -g team-api-server@$_TEAM_API_SERVER_VERSION"

USER root
RUN groupadd -r team-api && useradd -r -m -g team-api team-api && \
    mkdir -p $APP_SYS_ROOT/team-api/team-api.18f.gov && \
    chown team-api:team-api $APP_SYS_ROOT/team-api/team-api.18f.gov
VOLUME ["$APP_SYS_ROOT/team-api/team-api.18f.gov"]
USER team-api
WORKDIR $APP_SYS_ROOT/team-api

EXPOSE 6000 6001
COPY ["entrypoint.sh", "entrypoint.sh"]
ENTRYPOINT ["./entrypoint.sh"]
CMD ["run-server"]
