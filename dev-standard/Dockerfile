FROM dev-base

ENV _GO_VERSION go1.6
RUN /bin/bash -l -c "gvm install $_GO_VERSION && gvm use $_GO_VERSION --default"

ENV _RUBY_VERSION 2.3.0
RUN /bin/bash -l -c \
    "rbenv install $_RUBY_VERSION && rbenv global $_RUBY_VERSION && \
    gem install bundler colorator"

ENV _PYTHON_VERSION 2.7.11
ENV _PYTHON3_VERSION 3.5.1
RUN /bin/bash -l -c \
    "pyenv install $_PYTHON_VERSION && pyenv install $_PYTHON3_VERSION && \
    pyenv global $_PYTHON_VERSION"

ENV _NODE_VERSION 5.7.0
RUN /bin/bash -l -c \
    "nvm install $_NODE_VERSION && nvm alias default $_NODE_VERSION"
