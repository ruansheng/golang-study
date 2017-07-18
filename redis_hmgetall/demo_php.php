<?php
 
function getData($key) {
	$redis = new Redis();
   	$redis->connect('127.0.0.1', 6379);
    return $redis->hgetall($key);
}

define('SIZE', 10);

$data = [];

for($i = 1; $i <= SIZE; $i++) {
	$tmps = getData('users'.$i);
	foreach($tmps as $index => $val) {
		$data[$index] = $val;
	}
}

var_dump($data);
echo 'data len:' . count($data);