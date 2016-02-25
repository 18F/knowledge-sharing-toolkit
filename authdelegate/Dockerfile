FROM dev-standard

RUN bash -l -c 'go get github.com/18F/authdelegate && \
    strip $GOPATH/bin/authdelegate'

RUN groupadd -r authdelegate && useradd -r -m -g authdelegate authdelegate
USER authdelegate
WORKDIR $APP_SYS_ROOT/authdelegate

EXPOSE 9000
COPY ["entrypoint.sh", "entrypoint.sh"]
ENTRYPOINT ["./entrypoint.sh"]
CMD ["run-server"]
