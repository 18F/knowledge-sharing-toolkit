#! /bin/bash

_GO_VERSION=go1.6
_RUBY_VERSION=2.3.0
_PYTHON_VERSION=3.5.1
_NODE_VERSION=5.7.0

. /etc/profile.d/gvm.sh
#gvm install go1.4.3 && gvm use go1.4.3 && \
gvm install $_GO_VERSION && gvm use $_GO_VERSION --default && \
#gvm uninstall go1.4.3

rbenv install $_RUBY_VERSION && \
    rbenv global $_RUBY_VERSION && gem install bundler colorator

pyenv install $_PYTHON_VERSION && \
    pyenv global $_PYTHON_VERSION && pip install --upgrade pip

. /etc/profile.d/nvm.sh && nvm install $_NODE_VERSION && \
    nvm alias default $_NODE_VERSION && npm install -g forever
