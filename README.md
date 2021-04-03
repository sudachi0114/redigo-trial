# redigo trial

Trial connecting and operating Redis from Go client.

## Install

### Install Redis

* Install on local

```shell
# mac
$ brew install redis

Warning: Treating redis as a formula. For the cask, use homebrew/cask/redis
==> Downloading https://homebrew.bintray.com/bottles/redis-6.2.1.catalina.bottle
==> Downloading from https://d29vzk4ow07wi7.cloudfront.net/81c2a841e0b19040e7ea9
######################################################################## 100.0%
==> Pouring redis-6.2.1.catalina.bottle.tar.gz
==> Caveats
To have launchd start redis now and restart at login:
  brew services start redis
Or, if you don't want/need a background service you can just run:
  redis-server /usr/local/etc/redis.conf
==> Summary
🍺  /usr/local/Cellar/redis/6.2.1: 13 files, 2.0MB

```

FIXME: local でやろうとしたら見知らぬエラーが出てしまった。
わからないことが多い中で、複数のことをやるのはむずかしいので、いったんペンディング。

発生したエラーだけメモ ⬇️

```
sudachi@DaiMac:redigo-trial (develop *)
(*'-') < make
go build -o bin/main .
./bin/main
Redigo trial.
Succesfly Connect to redis @ 127.0.0.1:6379
panic: MISCONF Redis is configured to save RDB snapshots, but it is currently not able to persist on disk. Commands that may modify the data set are disabled, because this instance is configured to report errors during writes if RDB snapshotting fails (stop-writes-on-bgsave-err or option). Please check the Redis logs for details about the RDB error.

goroutine 1 [running]:
main.Set(0x124056a, 0x4, 0x1240562, 0x4, 0x1286c80, 0xc000076000, 0x0, 0x0)
        /Users/sudachi/go/src/github.com/sudachi0114/redigo-trial/main.go:45 +0x18f
main.main()
        /Users/sudachi/go/src/github.com/sudachi0114/redigo-trial/main.go:15 +0x105
make: *** [run] Error 2
```


### Create redis container by docker

docker で立てる方式に切り替える

* コンテナ作成

```shell
$ docker run -d -p 6379:6379 --name myredis redis
```

* 接続

```
$ docker exec -it myredis bash

root@14ab4257e657:/data# redis-
redis-benchmark  redis-check-aof  redis-check-rdb  redis-cli        redis-sentinel   redis-server     

root@14ab4257e657:/data# redis-server 
37:C 03 Apr 2021 03:12:08.245 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
37:C 03 Apr 2021 03:12:08.245 # Redis version=6.0.9, bits=64, commit=00000000, modified=0, pid=37, just started
37:C 03 Apr 2021 03:12:08.245 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
37:M 03 Apr 2021 03:12:08.247 # Could not create server TCP listening socket *:6379: bind: Address already in use

root@14ab4257e657:/data# redis-cli 
127.0.0.1:6379> KEYS *
(empty array)
127.0.0.1:6379> 
127.0.0.1:6379> 
127.0.0.1:6379> exit
root@14ab4257e657:/data# exit
exit
```

* 実行 (SET と GET)

```shell
$ go run main.go
Redigo trial.
Succesfly Connect to redis @ 127.0.0.1:6379
OK
fuga
```

* 実行後

```
$ docker exec -it myredis bash
root@14ab4257e657:/data# redis-cli 
127.0.0.1:6379> KEYS *
1) "hoge"
127.0.0.1:6379> GET hoge
"fuga"
```


### Install Redigo

Redis client library for go

```shell
$ go get github.com/gomodule/redigo/redis
```

## Links
* [GoとRedisにおける簡単なチャットアプリケーション](https://medium.com/eureka-engineering/go-redis-application-28c8c793a652)
* [GoでRedisをかるーくいじってみた](https://qiita.com/akubi0w1/items/8701c05fe7186ceee632)
