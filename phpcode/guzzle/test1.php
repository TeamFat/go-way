<?php
//http://guzzle-cn.readthedocs.io/zh_CN/latest/quickstart.html#id6
//https://github.com/guzzle/guzzle
require 'vendor/autoload.php';
$client = new \GuzzleHttp\Client();
$res = $client->request('GET', 'https://api.github.com/repos/guzzle/guzzle');
echo $res->getStatusCode();
// 200
echo $res->getHeaderLine('content-type');
// 'application/json; charset=utf8'
echo $res->getBody();
// '{"id": 1420053, "name": "guzzle", ...}'
//=================================================================
// Send an asynchronous request.
$request = new \GuzzleHttp\Psr7\Request('GET', 'http://httpbin.org');
$promise = $client->sendAsync($request)->then(function ($response) {
    echo 'I completed! ' . $response->getBody();
});
$promise->wait();
//================================================================
use GuzzleHttp\Client;
use GuzzleHttp\Promise;

$client = new Client(['base_uri' => 'http://httpbin.org/']);

// Initiate each request but do not block
$promises = [
    'image' => $client->getAsync('/image'),
    'png'   => $client->getAsync('/image/png'),
    'jpeg'  => $client->getAsync('/image/jpeg'),
    'webp'  => $client->getAsync('/image/webp')
];

// Wait on all of the requests to complete.
$results = Promise\unwrap($promises);

// You can access each result using the key provided to the unwrap
// function.
echo $results['image']->getHeader('Content-Length')[0];
echo $results['png']->getHeader('Content-Length')[0];