<?php
function getData($key, $port) {
	$redis = new Redis();
   	$redis->connect('47.94.226.123', $port);
	$redis->auth("asdfesdgrrdfgedfedsd");
    return $redis->hgetall($key);
}

define('SIZE', 10);

$redis_port = [
	6380,6381,6382,6383,6384,6385,6386,6387,6388,6389
];

$data = [];

for($i = 0; $i < SIZE; $i++) {
	$port = $redis_port[$i];
	$tmps = getData('users'.$i, $port);
	foreach($tmps as $index => $val) {
		$data[$index] = $val;
	}
}

var_dump($data);
echo 'data len:' . count($data);