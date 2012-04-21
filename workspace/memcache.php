<?php

$cache = new Memcache;
$key = 'cache_test';
$cache->connect('localhost', 11211);
// $cache->set($key, 'aaaa');
var_dump($cache->get($key));
