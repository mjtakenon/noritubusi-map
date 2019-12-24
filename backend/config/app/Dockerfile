FROM golang:1.13.0 
WORKDIR $GOPATH/src/noritubusi-map/backend/app

COPY . .

# Locale を ja_JP.UTF-8 にする (日本語対応)
RUN apt-get update && \
  apt-get install -y locales && \
  rm -rf /var/lib/apt/lists/* && \
  echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
  locale-gen ja_JP.UTF-8
ENV LC_ALL ja_JP.UTF-8

# To check mysql's healthceck
RUN apt-get update && apt-get install -y default-mysql-client

CMD ["./delayed_start"]
