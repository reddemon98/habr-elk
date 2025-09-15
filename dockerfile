FROM fluent/fluentd:v1.19-debian

USER root

# Устанавливаем плагины для Fluentd
RUN gem install fluent-plugin-opensearch fluent-plugin-gelf --no-document

# Устанавливаем netcat (версия OpenBSD)
RUN apt-get update && apt-get install -y netcat-openbsd && rm -rf /var/lib/apt/lists/*

# Запуск Fluentd с конфигом и плагинами
CMD ["fluentd", "-c", "/fluentd/etc/fluent.conf", "-p", "/fluentd/plugins"]
