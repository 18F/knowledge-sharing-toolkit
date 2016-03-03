FROM dev-standard

USER root
RUN groupadd -r pages && useradd -r -m -g pages pages && \
    mkdir -p $APP_SYS_ROOT/pages/bundle/cache \
      $APP_SYS_ROOT/pages/repos/pages.18f.gov \
      $APP_SYS_ROOT/pages/repos/pages-staging.18f.gov \
      $APP_SYS_ROOT/pages/repos/pages-internal.18f.gov \
      $APP_SYS_ROOT/pages/repos/pages-releases.18f.gov \
      $APP_SYS_ROOT/pages/sites/pages.18f.gov \
      $APP_SYS_ROOT/pages/sites/pages-staging.18f.gov \
      $APP_SYS_ROOT/pages/sites/pages-internal.18f.gov \
      $APP_SYS_ROOT/pages/sites/pages-releases.18f.gov && \
    chown pages:pages $APP_SYS_ROOT/pages/bundle/cache \
      $APP_SYS_ROOT/pages/repos/pages.18f.gov \
      $APP_SYS_ROOT/pages/repos/pages-staging.18f.gov \
      $APP_SYS_ROOT/pages/repos/pages-internal.18f.gov \
      $APP_SYS_ROOT/pages/repos/pages-releases.18f.gov \
      $APP_SYS_ROOT/pages/sites/pages.18f.gov \
      $APP_SYS_ROOT/pages/sites/pages-staging.18f.gov \
      $APP_SYS_ROOT/pages/sites/pages-internal.18f.gov \
      $APP_SYS_ROOT/pages/sites/pages-releases.18f.gov
USER 18f
VOLUME $APP_SYS_ROOT/pages/repos $APP_SYS_ROOT/pages/sites \
  $APP_SYS_ROOT/pages/bundle/cache

ENV _PAGES_VERSION=0.3.2 \
    _JEKYLL_VERSION=3.1.2 \
    _AWSCLI_VERSION=1.10.8
RUN bash -l -c "npm install -g 18f-pages-server@$_PAGES_VERSION && \
    gem install jekyll -v $_JEKYLL_VERSION && \
    pip install awscli==$_AWSCLI_VERSION"

# Needed to avoid the dreaded:
# Liquid Exception: incompatible character encodings: US-ASCII and UTF-8
ENV LANG=en_US.UTF-8
USER root
RUN locale-gen en_US.UTF-8 && update-locale LANG=en_US.UTF-8

USER pages
WORKDIR $APP_SYS_ROOT/pages
EXPOSE 5000

COPY ["entrypoint.sh", "entrypoint.sh"]
ENTRYPOINT ["./entrypoint.sh"]
CMD ["run-server"]
