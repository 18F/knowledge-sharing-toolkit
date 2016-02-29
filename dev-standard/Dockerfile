FROM dev-base

ENV _GO_VERSION=go1.6 \
    _RUBY_VERSION=2.3.0 \
    _PYTHON_VERSION=3.5.1 \
    _NODE_VERSION=5.7.0

RUN source $_GVM_PROFILE && gvm install go1.4.3 && gvm use go1.4.3 && \
    gvm install $_GO_VERSION && gvm use $_GO_VERSION --default && \
    gvm uninstall go1.4.3

RUN source $_RBENV_PROFILE && rbenv install $_RUBY_VERSION && \
    rbenv global $_RUBY_VERSION && gem install bundler colorator

RUN source $_PYENV_PROFILE && pyenv install $_PYTHON_VERSION && \
    pyenv global $_PYTHON_VERSION && pip install --upgrade pip

RUN source $_NVM_PROFILE && nvm install $_NODE_VERSION && \
    nvm alias default $_NODE_VERSION
