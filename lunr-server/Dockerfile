FROM dev-standard

ENV _LUNR_SERVER_VERSION 0.1.3
RUN bash -l -c "npm install -g lunr-server@$_LUNR_SERVER_VERSION"

RUN groupadd -r lunr-server && useradd -r -m -g lunr-server lunr-server
USER lunr-server
WORKDIR $APP_SYS_ROOT/lunr-server

EXPOSE 7777
COPY ["entrypoint.sh", "entrypoint.sh"]
ENTRYPOINT ["./entrypoint.sh"]
CMD ["run-server"]
