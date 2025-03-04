# Shortplz

Telegram Bot that does articles summarization. Powered by [Yandex API](300.ya.ru). Requires Yandex API token.

## `env`

`BOT_TOKEN` is a telegram bot token env var needs to be provided in `.env` file.

## Builing

- Linux x86_64

```sh
make build PLATFORM=linux
```

- OS X Arm64

```sh
make build PLATFORM=darwin
```

- Windows x64 (not tested)

```sh
make build PLATFORM=win64
```
