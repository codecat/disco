FROM alpine:latest

RUN apk add sudo zsh curl git vim tmux

RUN adduser -D -u 1000 -s /bin/zsh developer
RUN echo 'developer ALL=(ALL) NOPASSWD:ALL' > /etc/sudoers.d/developer
USER 1000

WORKDIR /home/developer
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

WORKDIR /src
CMD /bin/zsh
