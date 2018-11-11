conoha-api-go-client [日本語](README-ja.md)
=========

[![Build Status](https://travis-ci.org/is2ei/conoha-api-go-client.svg?branch=master&style=flat-square)][travis]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[travis]: https://travis-ci.org/is2ei/conoha-api-go-client
[godocs]: https://godoc.org/github.com/is2ei/conoha-api-go-client/conoha

conoha-api-go-client is a client written in golang to consume ConoHa API.

It also provides an handy CLI to easily interact with ConoHa API.

- [CLI](#cli)
- [lib](#lib)
- [Supported APIs](https://github.com/is2ei/conoha-api-go-client/wiki/Supported-APIs)

## CLI

### Install

Download latest binary from [Releases](https://github.com/is2ei/conoha-api-go-client/releases)

### Set conf file

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

### Get access token

```
$ conoha token
xxxxxtokenwillbeherexxxxx
```

### Get compute images

```
$ conoha compute images
name[vmi-mattermost-5.4.0-centos-7.5-amd64-20gb],id[e0347fa6-3b8b-41d7-907d-6ad39448306b]
name[vmi-mattermost-5.4.0-centos-7.5-amd64],id[98bab83e-4209-4034-9b33-f48fb6a2edbd]
name[vmi-isucon8-qualify-centos-7.5-amd64-20gb],id[ce76b3f6-7ae7-413b-8a72-260723d5597e]
...
```

### Get compute plans

```
$ conoha compute flavors
name[g-2gb],id[294639c7-72ba-43a5-8ff2-513c8995b869]
name[g-16gb],id[3aa001cd-95b6-46c9-a91e-e62d6f7f06a3]
name[g-4gb],id[62e8fb4b-6a26-46cd-be13-e5bbf5614d15]
...
```

### Get keypairs

```
$ conoha compute keypairs
name[abcdeasr]
name[asdfalskejr]
name[sajel;kjk]
...
```

### Add server

```
$ conoha compute server add \
-i [image id] \
-f [flavor id] \
-k [keypair name]
id[xxxxxxxxxxxxxxxxxxxx]
```

### Get servers

```
$ conoha compute servers
name[xxxx],id[xxxxxxxxxxxxxxxxxxxx]
name[xxxx],id[xxxxxxxxxxxxxxxxxxxx]
name[xxxx],id[xxxxxxxxxxxxxxxxxxxx]
...
```

### Delete servers

```
$ conoha compute server delete \
-i [server id]
```

## lib

See [GoDoc](https://godoc.org/github.com/is2ei/conoha-api-go-client/conoha)