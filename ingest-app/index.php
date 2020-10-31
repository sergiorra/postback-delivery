<?php

$redis = new Redis();

try {
    $redis->connect("redis", 6379);
} catch (RedisException $ex) {
    print_r($ex);
}

echo "Connected to Redis\n";

$json = file_get_contents("php://input");
$decoded = json_decode($json, true);

if ($decoded['data'] && $decoded['endpoint']) {
    $encoded = json_encode($decoded);
    $pushed = $redis->rPush("data", $encoded);
    if (!$pushed) {
        print_r("Error pushing data to Redis\n");
    } else {
        echo "Pushed to Redis\n";
    }
}

?>