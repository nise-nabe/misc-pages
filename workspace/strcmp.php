<?php

var_dump((bool) (0 < 'test')); # false
var_dump((bool) (0 < '1'));    # true
var_dump((bool) (0 < '-1'));   # false


