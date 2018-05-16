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
                'Authorization' => $conf['Authorization']]]);
$params['json'] = ['list'=>$data];
$url = $conf['server_host'] . 'api/key/store-many';
try{
    $res = $client->post($url,$params);
}catch (Exception $e){
    echo $e->getMessage();
}

echo '导入成功';


