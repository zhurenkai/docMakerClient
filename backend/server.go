package main

import(
  "net/http"
    "net/url"
    "net/http/httputil"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "bytes"
    "strings"
    "mime/multipart"
    "math/rand"
    "time"
    "io"
    "os"
)

type server struct {
    remoteServerHost string
    port int
}

type config struct {
    RemoteServerHost string `json:"remote_server_host"`
    ClientID         string `json:"client_id"`
    ClientSecret     string `json:"client_secret"`
    Port             int    `json:"port"`
    Db               struct {
        Host     string `json:"host"`
        Port     int    `json:"port"`
        Database string `json:"database"`
        Username string `json:"username"`
        Password string `json:"password"`
    } `json:"db"`
}

func main() {

    server := server{}
    server.serve()

}

func (s *server) serve (){
    s.loadConfig()
    mux := http.NewServeMux()
    hFile := http.FileServer(http.Dir(`../frontend/dist`))
    mux.HandleFunc(`/api/`,s.apiProxy)
    mux.HandleFunc(`/client-api/`,requestProxy)
    mux.Handle(`/`,hFile)
    fmt.Sprintln(`server started on http://localhost:%d`,s.port)
    addr := fmt.Sprintf(`:%d`,s.port)
    http.ListenAndServe(addr,mux)


}

func (s *server)loadConfig()  {
    conf,err := ioutil.ReadFile(`config.json`)
    if err !=nil  {
        fmt.Println(err)
    }
    var c config
    err = json.Unmarshal(conf,&c)
    if err !=nil  {
        fmt.Println(err)
    }
    s.remoteServerHost = c.RemoteServerHost
    s.port = c.Port
}


func (s *server)apiProxy(w http.ResponseWriter,r *http.Request) {
    remoteServerHost := `http://` + s.remoteServerHost
    remote,err := url.Parse(remoteServerHost)
    if err !=nil {
       fmt.Println(err)
    }

    proxy :=httputil.NewSingleHostReverseProxy(remote)
    // 重写rui
    b := []byte(r.RequestURI)
    r.RequestURI = string(b[4:])
    r.URL = &url.URL{Path:r.RequestURI}
    // 重设host否则就是反向代理，会将原来的请求带过去导致
    r.Host = s.remoteServerHost
    proxy.ServeHTTP(w, r)
}

func requestProxy(w http.ResponseWriter,r *http.Request) {
    ct :=r.Header.Get("Content-Type")
    var reqHeaders,reqMethod,reqHost,reqUri string
    var req *http.Request
    var err error
    var client http.Client
    if ct ==`application/json` {
        jsonInput,_ := ioutil.ReadAll(r.Body)
       // fmt.Println(jsonInput)
        var jsonContainer map[string]string
        json.Unmarshal(jsonInput,&jsonContainer)
       // fmt.Println(jsonContainer,err)
        reqHeaders = jsonContainer[`request_headers`]
        reqMethod = jsonContainer[`request_method`]
        reqHost = jsonContainer[`request_host`]
        reqUri = jsonContainer[`request_uri`]
        jsonData := jsonContainer[`json_input`]
        data :=bytes.NewReader([]byte(jsonData))
        req,err = http.NewRequest(reqMethod,reqHost + reqUri,data)
    } else {
        r.ParseMultipartForm(1048576)
        reqHeaders = strings.Join(r.Form[`request_headers`],``)
        reqMethod = strings.Join(r.Form[`request_method`],``)
        reqHost = strings.Join(r.Form[`request_host`],``)
        reqUri = strings.Join(r.Form[`request_uri`],``)
        data := &bytes.Buffer{}
        writer := multipart.NewWriter(data)
        for k,v := range r.Form {
            if k ==`request_headers` || k == `request_method` || k == `request_host` || k == `request_uri` {
                continue
            }
           writer.WriteField(k,strings.Join(v,``))
        }
        if r.MultipartForm.File !=nil  {
            rand.Seed(time.Now().UnixNano())
            // 异步文件处理
            ch := make(chan string)
            tmpFileNum := len(r.MultipartForm.File)
            // 文件写入表单
            for field := range r.MultipartForm.File{
                go func(ch chan string, field string) {
                file, fileHeader, _   := r.FormFile(field)
                tmpFile := `tmp/` + fileHeader.Filename
                f,err := os.OpenFile(tmpFile,os.O_WRONLY|os.O_CREATE, 0666)
                if err !=nil {
                    fmt.Println(err)
                }
                io.Copy(f,file)
                f.Close()
                fileWriter,err := writer.CreateFormFile(field,tmpFile)
                    fh, err := os.Open(tmpFile)
                    defer fh.Close()
                io.Copy(fileWriter,fh)
                ch<-tmpFile
                }(ch,field)
            }
           for i:=0;i<tmpFileNum;i++{
               <-ch
           }

           // 删除文件
            go func() {
                os.RemoveAll(`tmp`)
                os.Mkdir(`tmp`,0777)
            }()
        }

        writer.Close()
        contentType := writer.FormDataContentType()
        writer.Close()
        req,err = http.NewRequest(reqMethod,reqHost + reqUri,data)
        req.Header.Set(`Content-Type`,contentType)
        req.Header.Set("Connection", "Keep-Alive")
    }


    // process header
    var headersContainer map[string]string
    json.Unmarshal([]byte(reqHeaders),&headersContainer)
    for k,v:= range headersContainer {
        req.Header.Add(k,v)
    }
    resp,err :=client.Do(req)
    bodyBytes,_ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    fmt.Fprint(w,string(bodyBytes))
    if err != nil {
        fmt.Println(err)  
    }

}
