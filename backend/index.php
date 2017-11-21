<?php
require './vendor/autoload.php';
$request = $_REQUEST;
$request_uri = trim($_SERVER['REQUEST_URI'], '/');
$request_uri_arr = explode('/', $request_uri);
$passer = $request_uri_arr[0];
$function = strstr($request_uri_arr[1], '?', true) ?: $request_uri_arr[1];
$passer = ucfirst($passer);
require_once  './Controller/' . $passer . '.php';
$passer = new $passer;
call_user_func([$passer, $function]);
