conoha-api-go-client [日本語](README-ja.md)
=========

[![Build Status](https://travis-ci.org/is2ei/conoha-api-go-client.svg?branch=master&style=flat-square)][travis]
[![Coverage Status](https://coveralls.io/repos/github/is2ei/conoha-api-go-client/badge.svg?branch=master)][coverall]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]
[![Go Report Card](https://goreportcard.com/badge/github.com/is2ei/conoha-api-go-client)][goreport]

[travis]: https://travis-ci.org/is2ei/conoha-api-go-client
[coverall]: https://coveralls.io/github/is2ei/conoha-api-go-client?branch=master
[godocs]: https://godoc.org/github.com/is2ei/conoha-api-go-client/conoha
[goreport]: https://goreportcard.com/report/github.com/is2ei/conoha-api-go-client

conoha-api-go-client is a client written in golang to consume ConoHa API.

It also provides an handy CLI to easily interact with ConoHa API.

- [CLI](#cli)
- [lib](#lib)
- [Supported APIs](https://github.com/is2ei/conoha-api-go-client/wiki/Supported-APIs)

## CLI

### Install

Download latest binary from [Releases](https://github.com/is2ei/conoha-api-go-client/releases)

### Help

```sh
$ conoha --help
Yet another ConoHa API client built by is2ei

Usage:
  conoha [flags]
  conoha [command]

Available Commands:
  account        Account commands
  compute        Compute commands
  database       Database commands
  dns            DNS commands
  help           Help about any command
  identity       Identity commands
  image          Image commands
  mail           Mail commands
  network        Network commands
  servers        Get servers
  servers-detail Get servers detail
  token          Get token

Flags:
  -h, --help   help for conoha

Use "conoha [command] --help" for more information about a command.
```

### Set conf file

```sh
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

```sh
$ conoha token
xxxxxxTOKENxxxxxx
```

```sh
$ conoha token -v
token[xxxxxxTOKENxxxxxx], expires[2018-11-12T13:55:42Z]
```

### Servers list

```sh
$ conoha servers
name[xxxxxx], id[xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx]
name[xxxxxx], id[xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx]
name[xxxxxx], id[xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx]
...
```

```sh
$ conoha servers-detail
name[xxxxxx], id[xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx], status[ACTIVE]
name[xxxxxx], id[xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx], status[ACTIVE]
name[xxxxxx], id[xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx], status[ACTIVE]
...
```

### Add server

```sh
$ conoha server add \
-i [image id] \
-f [flavor id] \
-k [keypair name]
id[xxxxxxxxxxxxxxxxxxxx]
```

### Delete servers

```sh
$ conoha server delete \
-i [server id]
```

### Get compute images

```sh
$ conoha compute images
name[vmi-mattermost-5.4.0-centos-7.5-amd64-20gb], id[e0347fa6-3b8b-41d7-907d-6ad39448306b]
name[vmi-mattermost-5.4.0-centos-7.5-amd64], id[98bab83e-4209-4034-9b33-f48fb6a2edbd]
name[vmi-isucon8-qualify-centos-7.5-amd64-20gb], id[ce76b3f6-7ae7-413b-8a72-260723d5597e]
...
```

```sh
$ conoha compute images | grep centos # find centos image
name[vmi-centos-6.10-i386, id[3d34a870-0dfd-4d13-ab09-85b6f7a634de
name[vmi-centos-6.10-i386-20gb, id[91cab6a8-c5e7-4118-90bd-fb2ba4bb14df
name[vmi-centos-6.10-amd64, id[a7466aef-8213-4fda-997d-81303a841860
...
```

### Get compute plans

```sh
$ conoha compute flavors
name[g-2gb],id[294639c7-72ba-43a5-8ff2-513c8995b869]
name[g-16gb],id[3aa001cd-95b6-46c9-a91e-e62d6f7f06a3]
name[g-4gb],id[62e8fb4b-6a26-46cd-be13-e5bbf5614d15]
...
```

### Get keypairs

```sh
$ conoha compute keypairs
name[abcdeasr]
name[asdfalskejr]
name[sajel;kjk]
...
```

## lib

See [GoDoc](https://godoc.org/github.com/is2ei/conoha-api-go-client/conoha)

## Contribute

See [CONTRIBUTING.md](CONTRIBUTING.md)