<?php
return [
    'server_host'=>'http://docmaker-server.randqun.com/',
    'Authorization'=>'Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImFlYjQ3NGRlMDA2NDUxNTllMjY0ZDFhMTQyMTRhZjM0ODRmNDY4MjE2ZjhjYTlkNDE1ODFmNzQxZDk1MmQ5YTQ5Y2YwMTM4Mjk1YzIxZjkwIn0.eyJhdWQiOiIyIiwianRpIjoiYWViNDc0ZGUwMDY0NTE1OWUyNjRkMWExNDIxNGFmMzQ4NGY0NjgyMTZmOGNhOWQ0MTU4MWY3NDFkOTUyZDlhNDljZjAxMzgyOTVjMjFmOTAiLCJpYXQiOjE1MjY0ODQ1OTcsIm5iZiI6MTUyNjQ4NDU5NywiZXhwIjoxNTU4MDIwNTk3LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.HNigBFFlhU6BNtB7eOpXAalj7py2HsebqUNCOYQP6GrxiAHC3TFRTipp5iuCTgAJovnynzsPWXIT38wFT6Rc1CTw-vCsVr_fih-It52Wzqicd8hZbxJLS4pnnwv_01Ss2PezZhHIIb_euZ-PPEK9L0fGZO98PvSSYJOcLBkg-oV7ZbTudewpQIk3JPGaBCqsz7RbYOrBbKVtaO-FcqLKO_XdRTCHggC5dBoQ_rXuT5GSQRNEVP2xD_7jNd03-vXZmtA1X5sjR_48JfGbzRWuO9OyA7qUyJEUFz_OLriWzPcMWo82FwAfDa_dp6wuW_NYXL_PvoSyhtecvuu6OyQmOjlreJJe3TYZcDj2ySRkFCo6OS9us3eBptRw1VPvP4riq5Vp__0CDiSNWQpw36hjKh1DA6bqaWIzRFR2owrRtAikKYnaHdrV3fnQtwLyNEJGS8B9zODNE8spayZZOL5mn9HfP5cqltTqz2ylQHOS-EnWm6YXn6oaxaYVXkMDCqvsRniXmRXEMx3IivtWRTZ7qIsnbRkSbfs01R6SA9Fugw13K3Er3QZ_OBq60SUU6dv8sT_Hngo94gM9yOB5dtbVq5uD7Olnkz2a6U3buJdv2YV3gQo2zJsT1fJ9jDobCVI5jtucWRUPl-82gump9WwMv9Pr7eEPnClXAjFqNJrXGO8',
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