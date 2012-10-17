さくらの VPS で resolv.conf を編集する方法
==========================================


とりあえず resolvconf というツールを用いる必要があるらしい

下記ファイルに追記

/etc/resolvconf/resolv.conf.d/base
```
search sakura.ne.jp
nameserver 210.188.224.10
nameserver 210.188.224.11
```


/etc/resolv.conf をアップデートする

```
$ sudo resolvconf -v 
```

