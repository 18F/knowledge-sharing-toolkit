FROM dev-standard

RUN bash -l -c 'go get github.com/18F/authdelegate && \
    strip $GOPATH/bin/authdelegate'

RUN groupadd -r authdelegate && useradd -r -m -g authdelegate authdelegate
USER authdelegate
WORKDIR $APP_SYS_ROOT/authdelegate

EXPOSE 9000
CMD ["/bin/bash", "-l", "-c", "authdelegate config/authdelegate-config.json"]
