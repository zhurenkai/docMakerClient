<?php
use GuzzleHttp\Client;
require './vendor/autoload.php';
$conf = require 'config/config.php';
$dbms = $conf['db'];     //数据库类型
$db_conf = $conf['db_connections'][$dbms];
$host = $db_conf['host']; //数据库主机名
$database = $db_conf['database'];      //数据库连接用户名
$username = $db_conf['username'];;          //对应的密码
$password = $db_conf['password'];;          //对应的密码
$dsn = "$dbms:host=$host;dbname=$database";
$option = [PDO::MYSQL_ATTR_INIT_COMMAND => "set names ".$db_conf['charset']];
try {
    $db = new PDO($dsn,$username,$password,$option);
    echo "连接数据库成功".PHP_EOL;
} catch (PDOException $e) {
    die ("Error!: " . $e->getMessage() . "<br/>");
}

$sql = 'SELECT COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT FROM COLUMNS WHERE COLUMN_COMMENT !="" ORDER BY COLUMN_COMMENT DESC ';
$res = $db->query($sql);
$data = [];
while($row = $res->fetch(PDO::FETCH_ASSOC)){

    // your rule
    $data[lcfirst($row['COLUMN_NAME'])] = [
        'key'=>lcfirst($row['COLUMN_NAME']),
        'statement'=>$row['COLUMN_COMMENT'],
        'type'=>$row['DATA_TYPE'],
        'project_id'=>1,
        'user_id'=>1,
        'weight'=>0
    ];
}
$client = new Client(['headers'=>['Content-Type'=>'application/json','Accept' => 'application/json',
                'Authorization' => 'Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImFiOGJjMmVhYTJlYWNiNWE0OTAwNWFjYWQ0OGVhNGEzZjZkMmM1YTc4Y2JiNzQ2ZDc0NzE4ZjNmZGIxNzUwYzExMDhlZmQwYWQ3MWVmY2Q1In0.eyJhdWQiOiIyIiwianRpIjoiYWI4YmMyZWFhMmVhY2I1YTQ5MDA1YWNhZDQ4ZWE0YTNmNmQyYzVhNzhjYmI3NDZkNzQ3MThmM2ZkYjE3NTBjMTEwOGVmZDBhZDcxZWZjZDUiLCJpYXQiOjE1MTA4MzY2NjcsIm5iZiI6MTUxMDgzNjY2NywiZXhwIjoxNTQyMzcyNjY3LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.CEPGqaYI9rGNG2pTAYCmdXkTaVK7hbk1_8HOCWbpF0RD-tGKYbu0Sjocd4CHNN_lQxk0IfZLjXeAVCFIqi4OF8v4-t9lKdD2bJRe8EZOJM9SvDFmwZuDbsTJEHZJJ4o0WRHQyvk56Z9ENY0epMaxzFAcP8l_lyT3eEc_QfZHk_Ml6F4JPoEe2kCCgciL9YTnfXiRtcymbDhHZu-JKt0bwF1dMzk84JqVPTL5R4URJDOpyJ_yKnCJ9tgeUjEgW1UHWOC4lQdkf0Q3n2zrrdOW0M1S0N1g1WMfLRKjQrCy4EAltec9JtGoBp3zurgE2WvejiR9GDWmoI91yJkCZCQjmOQcG2KPilLpjtOd2K4JsX5vPcjOT2qMGXB49rMN9EnOippxqw_6VzQPdTPYgAafr6D0wYLVHZBCPZdfkHrsOOs_-bMT9Fe9NZIKCdjW3HD5lcVqCaA7vYdILNZKs-FdKVcj4BlIzhm7dCJF5fbCQvSfJb5kbwDkjLs7PjKjmyoh8D2l0hh7OhLVbK0rT3Ec4MyzY83XwalrcIqQYQa_0F9BwjmVvxvQ8sSnkvIumafFJoiht86Y1T2rzUSO-G7iwMCSwLlId9vGAbf3ivER68yJkrhR-6omguFfooI5-ON4ezIJDtE1FZUDEhJcAqeUBPwtLpDt6coTXHgwk9KBHlo']]);
$params['json'] = ['list'=>$data];
$url = $conf['server_host'] . 'api/key/store-many';
try{
    $res = $client->post($url,$params);
}catch (Exception $e){
    echo $e->getMessage();
}

echo '导入成功';


