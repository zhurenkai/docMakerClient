<?php
return [
    'server_host'=>'http://doc-server.local.com/',
    'server_uri'=>[
        'get_token'=>'oauth/token',
        'key_statement'=>'api/key-statement',
        'get_project_hosts'=>'api/',
    ],
    'db'=>'mysql',
    'db_connections'=>[
        'mysql' => [
            'driver' => 'mysql',
            'host' => '127.0.0.1',
            'port' => '3306',
            'database' => 'information_schema',
            'username' => 'root',
            'password' => '123',
            'unix_socket' => '',
            'charset' => 'utf8mb4',
            'collation' => 'utf8mb4_unicode_ci',
            'prefix' => '',
            'strict' => true,
            'engine' => null,
        ],
    ]
];