## KECE

An Experimental distributed Key Value Store written in Go

### TODO
- Add Pub Sub feature
- Protocol ? :D
- Support multiple datatype to store (now `Kece` only support simple string)

### Usage
- Build binary from source
```shell
$ go get github.com/Bhinneka/kece

$ go install github.com/Bhinneka/kece/cmd

$ kece --version
```

- Run `kece` server

    if `port flag` is not present, `kece` will using `9000` as the default port
```shell
$ kece -port 8000
```

- Store simple data
    
    you can use either `nc` or `telnet` as the client
```shell
$ nc localhost 8000
$
$ SET 1 wuriyanto
$ +OK
$
$ SET *BJE* bhinneka
$ +OK
$
$ GET 1
$ wurianto
$
$ GET *BJE*
$ bhinneka
$
$ DEL 1
$ +OK
```

- Auth mechanism

    if you want to use `Auth` on your `kece server`, simply add `-auth your-server-password` when start your server
```shell
$ kece -port 8000 -auth my-secret
```

    send auth to server
```shell
$ AUTH my-secret
$ +OK
$
```


#

### Author
Wuriyanto https://github.com/wuriyanto48