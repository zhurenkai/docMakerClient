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
    }
    protected function response($data,$status = true)
    {
        $response = [
            'data'=>$data,
            'status'=>$status,
        ];
        echo json_encode($response);exit;
    }
}