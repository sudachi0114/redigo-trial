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


## Links
* [https://medium.com/eureka-engineering/go-redis-application-28c8c793a652](https://medium.com/eureka-engineering/go-redis-application-28c8c793a652)