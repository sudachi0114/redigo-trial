# redigo trial

Trial connecting and operating Redis from Go client.

## Install

### Install Redis

* Install on local

```shell
# mac
$ brew install redis
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

```shell
$ make compose/up/sh
```

* 使用例 (現在はおうむ返しするだけ)

```
# ... (略)

/app # make
go build -o bin/main .
./bin/main
(屮`･д･)屮 hoge
hoge
(屮`･д･)屮 fuga
fuga
(屮`･д･)屮 piyo
piyo
(屮`･д･)屮 nyan
nyan
(屮`･д･)屮 exit
exit
(屮`･д･)屮 .exit
/app # exit
sudachi@DaiMac:redigo-trial (develop *)
(*'-') < 
```

これをチャットとして使えるようにしていく。
[参考](https://medium.com/eureka-engineering/go-redis-application-28c8c793a652)を写経しながら

## Redis

* プロセス

```
$ redis-server
```

* cli client

```
$ redis-cli
```

* 登録されている KEY を取得

```
> KEYS *
```

### 文字列

* 登録されている KEY の Value を取得

```
> GET hoge
```

* `KEY: hoge, Value: fuga` を登録

```
> SET hoge fuga
```

SET のオプション (?)
<!-- conn.DoでSETを実行しRedisに対して値を書き込みます。 
SETはデータを格納するためのコマンドです。 -->
- `NX` オプション: 同じ key が存在しない場合のみ、保存する。
- `EX`オプション* 指定した秒数後にデータが消去される。`EX` オプションを使用してキーの `TTL` を設定できる。(今回は120秒に設定する)
- `redigo/redis` では、KEY がすでに存在する場合、 `“ok”` ではなく `nil` の値を返す。これを用いて接続を試みたユーザーが既にオンラインかどうかを判断しています。

> TTLとは
> TTL は Time to Live の略で、日本語では有効生存期間、あるいは単に生存時間ということがある。
> ユーザーが設定した時間内に戻ってキーをリセットしない限り、自動で削除されるよう、キーのEXPIRE(有効期限)を設定しましょう。

⬇️ で `KEY` に設定された TTL の残り時間を確認できる。

```
TTL {{key}}
```

また、`TTL` は `EXPIRE` でも設定できる

```
> SET {{key}} EX {{sec}}

> EXPIRE {{key}} {{sec}}
```

### セット

* セットに member を追加

```
> SADD key member [member ...]
```

* セットの member 一覧を取得

```
> SMEMBERS key
```

* セットの member 数を取得

```
> SCARD key
```

今回は「ユーザーリスト」にセットを使っている。
> ユーザー名を SADD コマンドを使い users というリストにデータを追加していきます。
> `val, err = conn.Do(“SADD”, “users”, userName)`
> users というリストで接続したことあるユーザーを保存していきます。key/value をいつ更新するかをリマインドさせるGoのTickerをtimeパッケージのNewTickerを使い設定しましょう。


## Links
* [GoとRedisにおける簡単なチャットアプリケーション](https://medium.com/eureka-engineering/go-redis-application-28c8c793a652)
* [GoでRedisをかるーくいじってみた](https://qiita.com/akubi0w1/items/8701c05fe7186ceee632)

### Redis
* [Redisの起動と停止](https://qiita.com/horiko/items/bc812a03c9e0566d6338)
* [Redisで発生したメモリ不足エラーの調査メモ](http://www.24w.jp/blog/?p=82)
* [Redis Command - TTL(指定したキーの有効期間を確認)](https://symfoware.blog.fc2.com/blog-entry-531.html)

### docker-compose
* [docker-composeでredis環境をつくる](https://qiita.com/uggds/items/5e4f8fee180d77c06ee1)
* [Docker(+ Docker-Compose) に Redis を入れる](https://qiita.com/bonkoturyu/items/5e7e743b359ce63767a2)

### docker-compose go <-> redis
* [[Docker] golangとredisのコンテナを繋いでみた](https://shamaton.orz.hm/blog/archives/310)
* [Docker環境下でGoとRedisでAPIを実装する (Part.1)](https://qiita.com/Morero/items/473bc26ce2200c6a6fc6) ← これあとで別途やるかも..
