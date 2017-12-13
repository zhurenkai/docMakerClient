<?php
namespace Controller;
class Controller
{
    public $in;
    public $files;
    public $config;

    public function __construct()
    {
        $this->config = include_once './config/config.php';
        $this->in = $_REQUEST;
        $this->files = $_FILES;
        // 处理json
        if($_SERVER['CONTENT_TYPE']=='application/json'){
            $json_param = file_get_contents('php://input');
            $array = $this->jsonParam($json_param);
            $this->in = $array['json'];
        }
    }
    protected function response($data,$status = true)
    {
        $response = [
            'data'=>$data,
            'status'=>$status,
        ];
        echo json_encode($response);exit;
    }
    private function jsonParam($json){
        $params['json'] = json_decode($json,1);
        return $params;
    }
}