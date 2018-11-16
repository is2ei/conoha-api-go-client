conoha-api-go-client [English](README.md)
=========

[![Build Status](https://travis-ci.org/is2ei/conoha-api-go-client.svg?branch=master&style=flat-square)][travis]
[![Coverage Status](https://coveralls.io/repos/github/is2ei/conoha-api-go-client/badge.svg?branch=master)][coverall]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]
[![Go Report Card](https://goreportcard.com/badge/github.com/is2ei/conoha-api-go-client)][goreport]

[travis]: https://travis-ci.org/is2ei/conoha-api-go-client
[coverall]: https://coveralls.io/github/is2ei/conoha-api-go-client?branch=master
[godocs]: https://godoc.org/github.com/is2ei/conoha-api-go-client/conoha
[goreport]: https://goreportcard.com/report/github.com/is2ei/conoha-api-go-client

conoha-api-go-client は ConoHa の API を使用するためのクライアントです。

ConoHa の API を簡単に使用するための CLI も同梱されています。

- [CLI](#cli)
- [lib](#ライブラリ)
- [サポートしている API](https://github.com/is2ei/conoha-api-go-client/wiki/Supported-APIs)

## CLI

### インストール

[Releases](https://github.com/is2ei/conoha-api-go-client/releases) から最新のバージョンをダウンロードしてください。

### 設定ファイルを設置

```
$ cat ~/.conoha-api-go-client.conf
service_url:
  identity: https://identity.xxxx.conoha.io
  account: https://account.xxxx.conoha.io
  compute: https://compute.xxxx.conoha.io
  block_storage: https://block-storage.xxxx.conoha.io
  image: https://image-service.xxxx.conoha.io
  network: https://networking.xxxx.conoha.io
  object_storage: https://object-storage.xxxx.conoha.io
  database: https://database-hosting.xxxx.conoha.io
  dns: https://dns-service.xxxx.conoha.io
  mail: https://mail-hosting.xxxx.conoha.io
user:
  username: YOUR_USERNAME
  password: YOUR_PASSWORD
  tenant_id: YOUR_TENANT_ID
```

### アクセストークンの取得

```
$ conoha token
xxxxxtokenwillbeherexxxxx
```

## ライブラリ

[GoDoc](https://godoc.org/github.com/is2ei/conoha-api-go-client/conoha) を参照してください。

## コントリビュート

[CONTRIBUTING.md](CONTRIBUTING.md) を参照してください。