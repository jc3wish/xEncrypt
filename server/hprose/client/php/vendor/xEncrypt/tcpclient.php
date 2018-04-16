<?php
require_once dirname(dirname(__FILE__))."/autoload.php";
use \Hprose\Future;
use \Hprose\Socket\Client;
$test = new Client("tcp://127.0.0.1:1036",false);
/*
$test->fullDuplex = true;
$var_dump = Future\wrap("var_dump");
Future\co(function() use ($test) {
    try {
        var_dump((yield $test->DoEncrypt("email1","mytestname")));

        var_dump((yield $test->DoDecrypt("email1","N1/u1DamOOOHzxOExyJHQA==")));
        
        var_dump((yield $test->hello("testname")));
        var_dump((yield $test->hello("yield world1")));
        var_dump((yield $test->hello("yield world2")));
        var_dump((yield $test->hello("yield world3")));
        var_dump((yield $test->hello("yield world4")));
        var_dump((yield $test->hello("yield world5")));
        var_dump((yield $test->hello("yield world6")));
        
    }
    catch (\Exception $e) {
        echo ($e);
    }
});
*/

var_dump( $test->hello("testname"));

var_dump( $test->DoEncrypt("email1","mytestname"));
var_dump( $test->DoEncrypt("email2","mytestname"));
var_dump( $test->DoEncrypt("email3","mytestname"));
var_dump( $test->DoEncrypt("email4","mytestname"));

$client = \Hprose\Client::create('tcp://127.0.0.1:1036', false);
$data = $client->hello("testname");
print_r($data);



//$data = $test->DoDecrypt("email1","N1/u1DamOOOHzxOExyJHQA==");
//print_r( $data);

/*
$var_dump($test->hello("async world1"));
$var_dump($test->hello("async world2"));
$var_dump($test->hello("async world3"));
$var_dump($test->hello("async world4"));
$var_dump($test->hello("async world5"));
$var_dump($test->hello("async world6"));
$test->hello("world1")
->then(function($result) use ($test) {
    var_dump($result);
    return $test->hello("world2");
})
->then(function($result) use ($test) {
    var_dump($result);
    return $test->hello("world3");
})
->then(function($result) use ($test) {
    var_dump($result);
    return $test->hello("world4");
})
->then(function($result) use ($test) {
    var_dump($result);
    return $test->hello("world5");
})
->then(function($result) use ($test) {
    var_dump($result);
    return $test->hello("world6");
})
->then(function($result) {
    var_dump($result);
});
*/