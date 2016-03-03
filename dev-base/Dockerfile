FROM ubuntu:14.04.4
RUN rm /bin/sh && ln -s /bin/bash /bin/sh

# LANG, locale-gen, and update-locale are necessary to avoid the Jekyll error:
# Liquid Exception: incompatible character encodings: US-ASCII and UTF-8
ENV APP_SYS_ROOT=/usr/local/18f \
    LANG=en_US.UTF-8

RUN locale-gen en_US.UTF-8 && update-locale LANG=en_US.UTF-8 && \
    apt-get install -y software-properties-common && \
    add-apt-repository ppa:git-core/ppa && \
    apt-get update && apt-get install -y \
    bison \
    build-essential \
    curl \
    git \
    libbz2-1.0 \
    libbz2-dev \
    libcurl4-openssl-dev \
    libffi-dev \
    libpcre3-dev \
    libreadline-dev \
    libsqlite3-dev \
    libssl-dev \
    libxml2-dev \
    libxslt1-dev \
    libyaml-dev \
    libzip-dev \
    libzip2 \
    openssl \
    sqlite3 \
    vim \
    xz-utils \
    zip \
    zlib1g-dev \
    && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install Ruby via rbenv
ENV _RBENV_ROOT="$APP_SYS_ROOT/rbenv" \
    _RBENV_VERSION=v1.0.0 \
    _RUBY_BUILD_VERSION=v20160228 \
    _RBENV_PROFILE=/etc/profile.d/rbenv.sh
RUN git clone https://github.com/rbenv/rbenv.git $_RBENV_ROOT && \
    cd $_RBENV_ROOT && git checkout $_RBENV_VERSION 2>/dev/null && \
    mkdir $_RBENV_ROOT/plugins && \
    export _RUBY_BUILD_ROOT="$_RBENV_ROOT/plugins/ruby-build" && \
    git clone https://github.com/rbenv/ruby-build.git $_RUBY_BUILD_ROOT && \
    cd $_RUBY_BUILD_ROOT && git checkout $_RUBY_BUILD_VERSION 2>/dev/null && \
    echo export RBENV_ROOT=$_RBENV_ROOT >> $_RBENV_PROFILE && \
    echo 'export PATH="$RBENV_ROOT/bin:$PATH"' >> $_RBENV_PROFILE && \
    echo 'eval "$(rbenv init -)"' >> $_RBENV_PROFILE && \
    chmod +x $_RBENV_PROFILE

# Instally Python via pyenv
ENV _PYENV_ROOT="$APP_SYS_ROOT/pyenv" \
    _PYENV_VERSION=v20160202 \
    _PYENV_PROFILE=/etc/profile.d/pyenv.sh
RUN git clone https://github.com/yyuu/pyenv.git $_PYENV_ROOT && \
    cd $_PYENV_ROOT && git checkout $_PYENV_VERSION 2>/dev/null && \
    echo export PYENV_ROOT=$_PYENV_ROOT >> $_PYENV_PROFILE && \
    echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> $_PYENV_PROFILE && \
    echo 'eval "$(pyenv init -)"' >> $_PYENV_PROFILE && \
    chmod +x $_PYENV_PROFILE

# Install Node.js via nvm
ENV _NVM_ROOT="$APP_SYS_ROOT/nvm" \
    _NVM_VERSION=v0.31.0 \
    _NVM_PROFILE=/etc/profile.d/nvm.sh
RUN git clone https://github.com/creationix/nvm.git $_NVM_ROOT && \
    cd $_NVM_ROOT && git checkout $_NVM_VERSION 2>/dev/null && \
    echo . $_NVM_ROOT/nvm.sh >> $_NVM_PROFILE && \
    chmod +x $_NVM_PROFILE

# Install Go via gvm
ENV _GVM_ROOT="$APP_SYS_ROOT/gvm" \
    _GVM_VERSION=25ea8ae158e2861c92e2b22c458e60840157832f \
    _GVM_PROFILE=/etc/profile.d/gvm.sh
RUN curl -L -O https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer && \
    GVM_NO_UPDATE_PROFILE=true bash ./gvm-installer $_GVM_VERSION $APP_SYS_ROOT && \
    echo . $_GVM_ROOT/scripts/gvm >> $_GVM_PROFILE && \
    chmod +x $_GVM_PROFILE

RUN groupadd -r 18f && useradd -r -m -g 18f 18f && mkdir -p $APP_SYS_ROOT && \
    chown -R 18f:18f $APP_SYS_ROOT
USER 18f
