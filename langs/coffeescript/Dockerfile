FROM alpine:3.20 AS builder

RUN apk add --no-cache curl icu-data npm

WORKDIR /usr

RUN curl -L https://github.com/jashkenas/coffeescript/tarball/2.7.0 \
  | tar xz --strip-components 1

RUN npm install --global coffeescript

FROM codegolf/lang-base

COPY --from=0 /bin/cat /bin/sh                      /bin/
COPY --from=0 /lib/ld-musl-x86_64.so.1              \
              /lib/libcrypto.so.3                   \
              /lib/libssl.so.3                      \
              /lib/libz.so.1                        /lib/
COPY --from=0 /usr/package.json                     /usr/
COPY --from=0 /usr/bin/coffee                       \
              /usr/bin/env                          \
              /usr/bin/node                         /usr/bin/
COPY --from=0 /usr/lib/libada.so.2                  \
              /usr/lib/libbrotlicommon.so.1         \
              /usr/lib/libbrotlidec.so.1            \
              /usr/lib/libbrotlienc.so.1            \
              /usr/lib/libcares.so.2                \
              /usr/lib/libgcc_s.so.1                \
              /usr/lib/libicudata.so.74             \
              /usr/lib/libicui18n.so.74             \
              /usr/lib/libicuuc.so.74               \
              /usr/lib/libnghttp2.so.14             \
              /usr/lib/libstdc++.so.6               \
              /usr/lib/libuv.so.1                   /usr/lib/
COPY --from=0 /usr/lib/coffeescript/coffeescript.js \
              /usr/lib/coffeescript/command.js      \
              /usr/lib/coffeescript/helpers.js      \
              /usr/lib/coffeescript/index.js        \
              /usr/lib/coffeescript/lexer.js        \
              /usr/lib/coffeescript/nodes.js        \
              /usr/lib/coffeescript/optparse.js     \
              /usr/lib/coffeescript/parser.js       \
              /usr/lib/coffeescript/register.js     \
              /usr/lib/coffeescript/rewriter.js     \
              /usr/lib/coffeescript/scope.js        \
              /usr/lib/coffeescript/sourcemap.js    /usr/lib/coffeescript/
COPY --from=0 /usr/share/icu/74.2/icudt74l.dat      /usr/share/icu/

COPY /coffeescript /usr/bin/

ENTRYPOINT ["coffeescript"]

CMD ["--version"]
