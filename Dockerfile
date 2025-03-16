# Imagem base do Go mais recente
FROM golang:1.23-bullseye

# Instalação de dependências
RUN apt-get update && apt-get install -y \
  wget \
  git \
  bash \
  locales \
  && rm -rf /var/lib/apt/lists/*

# Instalação do Oh My Bash
RUN bash -c "$(wget https://raw.githubusercontent.com/ohmybash/oh-my-bash/master/tools/install.sh -O -)"

# Configuração do tema power no Oh My Bash
RUN sed -i 's/OSH_THEME="font"/OSH_THEME="powerline"/g' ~/.bashrc

# Instalação e configuração da localidade en_US.UTF-8
RUN sed -i '/en_US.UTF-8/s/^# //g' /etc/locale.gen \
  && locale-gen en_US.UTF-8

ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8

# Criação do diretório .go-cache
RUN mkdir -p /go/pkg
RUN mkdir -p ~/.go-cache
RUN chmod -R 755 ~/.go-cache

# Configuração do diretório de trabalho
WORKDIR /monitor
# Comando padrão ao iniciar o container
CMD ["bash"]
