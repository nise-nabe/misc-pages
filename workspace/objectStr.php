<?php

class A {
  public function __toString() {
    return '支離滅裂';
  }
}

$a = new A();

echo $a; // 支離滅裂
echo "\n";
echo ((string) $a); // 支離滅裂
echo "\n";

var_dump($a); // object(A)#1 (0) {}
var_dump((string) $a); // string(12) "支離滅裂"

var_dump('支離滅裂' == $a); // bool(true)
var_dump('支離滅裂' === $a); // bool(false)
var_dump( ! '' ); // bool(true)
