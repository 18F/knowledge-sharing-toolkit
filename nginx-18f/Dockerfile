FROM dev-standard

ENV NGINX_VERSION 1.9.11
ENV ZLIB_VERSION 1.2.8
ENV PCRE_VERSION 8.38
ENV OPENSSL_VERSION 1.0.2f

RUN mkdir -p $APP_SYS_ROOT/src && cd $APP_SYS_ROOT/src \
    && sh -c "curl http://nginx.org/download/nginx-$NGINX_VERSION.tar.gz | gzip -dc | tar xf -" \
    && sh -c "curl http://zlib.net/zlib-$ZLIB_VERSION.tar.gz | gzip -dc | tar xf -" \
    && sh -c "curl ftp://ftp.csx.cam.ac.uk/pub/software/programming/pcre/pcre-$PCRE_VERSION.tar.gz | gzip -dc | tar xf -" \
    && sh -c "curl https://www.openssl.org/source/openssl-$OPENSSL_VERSION.tar.gz | gzip -dc | tar xf -"

RUN cd $APP_SYS_ROOT/src/nginx-$NGINX_VERSION \
    && ./configure --prefix=/usr/local/18f/nginx \
      --conf-path=/usr/local/18f/nginx/config/nginx.conf \
      --with-cc=/usr/bin/gcc \
      --with-cc-opt='-I/usr/local/18f/include' \
      --with-http_addition_module \
      --with-http_auth_request_module \
      --with-http_gzip_static_module \
      --with-http_secure_link_module \
      --with-http_ssl_module \
      --with-http_stub_status_module \
      --with-http_sub_module \
      --with-http_v2_module \
      --with-ipv6 \
      --with-ld-opt='-L/usr/local/18f/lib -Wl,-rpath=/usr/local/18f/lib' \
      --with-md5-asm \
      --with-openssl=../openssl-$OPENSSL_VERSION \
      --with-pcre=../pcre-$PCRE_VERSION \
      --with-sha1-asm \
      --with-zlib=../zlib-$ZLIB_VERSION \
      --without-http_browser_module \
      --without-http_geo_module \
      --without-http_memcached_module \
      --without-http_scgi_module \
      --without-http_userid_module \
      --without-http_uwsgi_module \
   && make install && rm -rf $APP_SYS_ROOT/src \
   && strip /usr/local/18f/nginx/sbin/nginx

RUN groupadd -r nginx && useradd -r -g nginx nginx

WORKDIR $APP_SYS_ROOT/nginx
EXPOSE 80 443
CMD ["sbin/nginx", "-g", "daemon off;" ]
