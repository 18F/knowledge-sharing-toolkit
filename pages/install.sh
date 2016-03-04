#! /bin/bash

APP_SYS_ROOT=/usr/local/18f

if [ ! -d $APP_SYS_ROOT/pages/repos ]; then
  mkdir -p $APP_SYS_ROOT/pages/repos
fi

if [ ! -d $APP_SYS_ROOT/pages/sites ]; then
  mkdir -p $APP_SYS_ROOT/pages/sites
fi

_PAGES_VERSION=0.3.4 \
_JEKYLL_VERSION=3.1.2 \
_AWSCLI_VERSION=1.10.8
npm install -g 18f-pages-server@$_PAGES_VERSION && \
    gem install jekyll -v $_JEKYLL_VERSION && \
    pip install awscli==$_AWSCLI_VERSION
