FROM codecatt/disco:base
RUN sudo apk add npm nodejs
RUN sudo npm install -g pnpm
