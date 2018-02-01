<?php

use Controller\Controller;
use GuzzleHttp\Client as GClient;

class client extends Controller
{
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * 转发浏览器来的请求到本地的服务
     * 支持普通参数 表单文件 json
     *
     */
    public function request()
    {
        $request_method = $this->in['request_method'];
        if(!$host = $this->in['request_host']){
            $this->response('缺少host',false);
        }
        if(!$uri = $this->in['request_uri']){
            $this->response('缺少uri',false);
        }
        $url = $host.$uri;
        unset($this->in['request_uri']);
        unset($this->in['request_host']);
        unset($this->in['request_method']);
        $headers_str = $this->in['request_headers'] ?? null;
        if($headers_str){
            $headers = json_decode($headers_str,1);
        }
        $client = new GClient;
        // 处理带图片的
        if ($files = $_FILES) {
           $params = $this->multipartParam($this->in,$files);
        }
        // json
        if(isset($this->in['json_input'])){
            $params['json'] = json_decode($this->in['json_input'],1);
        }
        // 默认
        if(!isset($params)){
            $params['form_params'] = $this->in;
        }
        if(isset($headers)){
            $params['headers'] = $headers;
        }
            $res = $client->request($request_method, $url, $params);
        echo  $res->getBody() ;
    }

    private function getHeader($sever)
    {
        $header = [];
        foreach ($sever as $k => $v) {
            if (strpos($k, 'HTTP_') !== false) {
                $header[str_replace('HTTP_', '', $k)] = $v;
            }
        }
        unset($header['HOST']);
        unset($header['CACHE_CONTROL']);
        unset($header['CONNECTION']);
        return $header;
    }



    private function multipartParam($files,$in)
    {
        foreach ($files as $p_name => $file) {
            if (is_array($file['name'])) {
                foreach ($file['name'] as $index => $v) {
                    $params['multipart'][] = [
                        'name' => $p_name . '[]',
                        'contents' => fopen($file['tmp_name'][$index], 'r'),
                        'filename' => $v
                    ];
                }
            } else {
                $params['multipart'][] = [
                    'name' => $p_name,
                    'contents' => fopen($file['tmp_name'], 'r'),
                    'filename' => $file['name']
                ];
            }
        }

        foreach ($in as $name => $value) {
            $params['multipart'][] = [
                'name' => $name,
                'contents' => $value
            ];
        }
        return $params;
    }
}