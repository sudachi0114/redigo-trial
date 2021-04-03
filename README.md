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


### Create redis container by docker-compose

`docker-compose` で Redis コンテナを用意する。
Go のサービスと合わせて使うため。

* コンテナ作成

```shell
$ make compose/up
```

* 接続

```
$ docker-compose exec redis bash
```


### Install Redigo

Redis client library for go

```shell
$ go get github.com/gomodule/redigo/redis
```

## execution

* コンテナを作る

```
$ make compose/up
```

* 実行

```
$ make
```

## Links
* [GoとRedisにおける簡単なチャットアプリケーション](https://medium.com/eureka-engineering/go-redis-application-28c8c793a652)
* [GoでRedisをかるーくいじってみた](https://qiita.com/akubi0w1/items/8701c05fe7186ceee632)

### Redis
* [Redisの起動と停止](https://qiita.com/horiko/items/bc812a03c9e0566d6338)
* [Redisで発生したメモリ不足エラーの調査メモ](http://www.24w.jp/blog/?p=82)
