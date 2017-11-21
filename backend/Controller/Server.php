<?php
use Controller\Controller;
use GuzzleHttp\Client;

class Server extends Controller
{
    private $server_host;
    private $access_token;
    private $request_header;

    public function __construct()
    {
        parent::__construct();
        $this->server_host = $this->config['server_host'];
        $this->server_uri = $this->config['server_uri'];
        session_start();
        if (isset($_SESSION['auth'])) {
            $auth = $_SESSION['auth'];
            $this->request_header = [
                'Accept' => 'application/json',
                'Authorization' => $auth->token_type . ' ' . $auth->access_token
            ];
        }


    }

    public function getAuth()
    {
        $access_token_file = CACHE_DIR . ACCESS_TOKEN_CACHE_FILE;

        $params['form_params'] = [
            'grant_type' => 'password',
            'username' => $this->in['username'],
            'password' => $this->in['password'],
            'client_id' => 3,
            'client_secret' => 'rjTVYer2N4jZetiERvjw3WAxk3WTnZSHxRuzrBgl'];
        $client = new Client();
        $res = $client->post($this->server_host . $this->config['server_uri']['get_token'], $params);
        $res_obj = json_decode($res->getBody());
        file_put_contents($access_token_file, $res_obj->token_type . ' ' . $res_obj->access_token);
        $refresh_token_file = CACHE_DIR . REFRESH_TOKEN_CACHE_FILE;
        file_put_contents($refresh_token_file, $res_obj->refresh_token);
        $this->access_token = $res_obj->access_token;
        //session_start();
        $_SESSION['auth'] = $res_obj;
        print_r($res_obj);

    }

    public function getKeyState()
    {
        $key = trim($this->in['key']);
        $project_id = trim($this->in['project_id']);
        $client = $this->authClient();
        //var_dump($_SESSION);exit;
        // var_dump($this->request_header);
        $query = [
            'key' => $key,
            'project_id' => $project_id
        ];
        $res = $client->get($this->server_host . $this->config['server_uri']['key_statement'], ['query' => $query]);
        echo $res->getBody();
    }

    private function authClient()
    {
        return new Client(['headers' => $this->request_header]);
    }

    public function getProjectHosts()
    {
//        $res = $this->authClient()->get($this->server_host.$this->server_uri[''])
    }
}