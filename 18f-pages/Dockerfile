FROM dev-standard

ENV _PAGES_VERSION 0.2.1
ENV _JEKYLL_VERSION 3.1.2
RUN bash -l -c "npm install -g 18f-pages-server@$_PAGES_VERSION && \
    gem install jekyll -v $_JEKYLL_VERSION"

RUN groupadd -r 18f-pages && useradd -r -m -g 18f-pages 18f-pages && \
    mkdir $APP_SYS_ROOT/pages && chown 18f-pages:18f-pages $APP_SYS_ROOT/pages
USER 18f-pages
WORKDIR $APP_SYS_ROOT/pages
RUN mkdir conf repos sites
VOLUME $APP_SYS_ROOT/pages/sites
COPY ["./pages-config.json", "conf/pages-config.json"]
EXPOSE 5000
CMD ["/bin/bash", "-l", "-c", "18f-pages conf/pages-config.json"]
