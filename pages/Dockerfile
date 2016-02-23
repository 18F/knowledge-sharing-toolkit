FROM dev-standard

ENV _PAGES_VERSION 0.2.1
ENV _JEKYLL_VERSION 3.1.2
RUN bash -l -c "npm install -g 18f-pages-server@$_PAGES_VERSION && \
    gem install jekyll -v $_JEKYLL_VERSION"

RUN groupadd -r pages && useradd -r -m -g pages pages && \
    mkdir -p $APP_SYS_ROOT/pages/repos && \
    chown pages:pages $APP_SYS_ROOT/pages/repos
VOLUME $APP_SYS_ROOT/pages/repos
USER pages
WORKDIR $APP_SYS_ROOT/pages

EXPOSE 5000
CMD ["/bin/bash", "-l", "-c", "18f-pages pages-config.json"]
