FROM alpine:edge

LABEL org.opencontainers.image.source=https://github.com/codecat/disco

RUN apk add sudo fish curl git vim helix helix-tree-sitter-vendor tmux openssh file ripgrep

RUN adduser -D -u 1000 -s /usr/bin/fish developer
RUN echo 'developer ALL=(ALL) NOPASSWD:ALL' > /etc/sudoers.d/developer
USER 1000

WORKDIR /src
CMD /usr/bin/fish
