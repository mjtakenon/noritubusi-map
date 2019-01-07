FROM mysql:5.7

# Locale を ja_JP.UTF-8 にする (日本語対応)
RUN apt-get update && \
  apt-get install -y locales && \
  rm -rf /var/lib/apt/lists/* && \
  echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
  locale-gen ja_JP.UTF-8
ENV LC_ALL ja_JP.UTF-8

# 指定の場所にログを記録するファイルを作る
RUN touch /var/log/mysql/mysqld.log 
RUN chown -R mysql:mysql /var/log/mysql
