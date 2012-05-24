java Too many open Files
========================

対処法
------

(http://d.hatena.ne.jp/necoyama3/20110505/1304584584)

下記コマンドでファイルディスクリプタの数を確認

    $ ulimit -a | grep "open files"

下記コマンドでファイルディスクリプタの数を変更

    $ ulimit -n ファイルディスクリプタの数
