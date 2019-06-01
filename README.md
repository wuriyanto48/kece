<div align="center">

[<img src="./assets/logo-kece-readme.png" width="300">](https://github.com/wuriyanto48/kece)
<br/><br/>
[![Build Status](https://travis-ci.org/wuriyanto48/kece.svg?branch=master)](https://travis-ci.org/wuriyanto48/kece)
</div>

### What is kece?
An Experimental distributed Key Value Store written in Go


### TODO
- Add Pub Sub feature
- Protocol ? :D
- Support multiple datatype to store (now `Kece` only support simple string)

### Usage
- <b>Build binary from source</b>
```shell
$ go get github.com/wuriyanto48/kece

$ go install github.com/wuriyanto48/kece/cmd

$ kece --version
```

- <b>Run `kece` server</b>

    if `port flag` is not present, `kece` will using `9000` as the default port
```shell
$ kece -port 8000
 _  __ _____  ______  _____
| |/ /| |__| |   ___|| |__| |
| |\ \| |___ |  |    | |___
|____________|_____________**%**

log -> kece server listen on port : 8000
```

- There are two type of data structure for store data, `HashMap` and `Binary Tree` (default using `HashMap`). For choose data structure type, add flag `-ds`.
```shell
$ kece -port 8000 -ds bt
$ kece -port 8000 -ds hashmap
```

- <b>Store simple data</b>

    you can use either `nc` or `telnet` as the client
```shell
$ nc localhost 8000
$
$ SET 1 wuriyanto
$ +OK
$
$ SET *BJE* hello
$ +OK
$
$ GET 1
$ wurianto
$
$ GET *BJE*
$ hello
$
$ DEL 1
$ +OK
$
$ SET cache "this is cache value with lifetime 20 seconds" 20
$ +OK
```

- <b>Auth mechanism</b>

    if you want to use `Auth` on your `kece server`, simply add `-auth your-server-password` when start your server
```shell
$ kece -port 8000 -auth my-secret
 _  __ _____  ______  _____
| |/ /| |__| |   ___|| |__| |
| |\ \| |___ |  |    | |___
|____________|_____________**%**

log -> kece server listen on port : 8000
```

    send auth to server
```shell
$ AUTH my-secret
$ +OK
$
```

- <b>Access KECE from code</b>

    follow this repository https://github.com/Bhinneka/kece-client-examples to see example how to access `kece` from specific language


#

### Author
Wuriyanto https://github.com/wuriyanto48


### Contibutions PR
Before creating PR make sure your PR is passed.
Use the linter first, then commit and push

```
$ make lint-prepare

$ make lint

$ make test
```
