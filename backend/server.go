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
    "path/filepath"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"

    "crypto/md5"
)

type server struct {
    config config
    path string

}

type config struct {
    RemoteServerHost string `json:"remote_server_host"`
    ClientID         int `json:"client_id"`
    ClientSecret     string `json:"client_secret"`
    Port             int    `json:"port"`
}

type dbConfig struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Databases string `json:"databases"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func main() {
    server := NewServer()
    server.serve()
}

func (s *server) serve (){
    mux := http.NewServeMux()
    hFile := http.FileServer(http.Dir(filepath.Dir(s.path) + `/frontend/dist`))
    mux.HandleFunc(`/api/`,s.apiProxy)
    mux.HandleFunc(`/client-api/`,s.requestProxy)
    mux.HandleFunc(`/api/api/import-db-comments`,s.importDBComments)
    mux.HandleFunc(`/client-info`,s.clientInfo)
    mux.Handle(`/`,hFile)
    fmt.Print(`server started on http://localhost:`,s.config.Port)
    addr := fmt.Sprintf(`:%d`,s.config.Port)
    http.ListenAndServe(addr,mux)
}

func NewServer() *server{
    s := new(server)
    s.getPath()
    s.loadConfig()
    return s
}

func (s *server)getPath()  {
    dir,_ :=filepath.Abs(filepath.Dir(os.Args[0]))
    s.path = dir
   // s.path = `/home/akon/project/docMakerClient/backend`
}

func (s *server)loadConfig()  {
    conf,err := ioutil.ReadFile(s.path + `/config.json`)
    if err !=nil  {
        fmt.Println(err)
    }
    var c config
    err = json.Unmarshal(conf,&c)
    if err !=nil  {
        fmt.Println(err)
    }
    s.config = c
}

// 转发对后端api的请求
func (s *server)apiProxy(w http.ResponseWriter,r *http.Request) {
    remoteServerHost := `http://` + s.config.RemoteServerHost
    remote,err := url.Parse(remoteServerHost)
    if err !=nil {
       fmt.Println(err)
    }
    proxy :=httputil.NewSingleHostReverseProxy(remote)
    // 重写rui
    b := []byte(r.URL.Path)
    newUrl := string(b[4:])
    r.RequestURI = newUrl
    r.URL = &url.URL{Path:newUrl,RawQuery:r.URL.RawQuery}
    // 重设host,否则就是反向代理，会将原来的请求带过去
    r.Host = s.config.RemoteServerHost
    proxy.ServeHTTP(w, r)
}

// 转发用户的请求
func (s *server)requestProxy(w http.ResponseWriter,r *http.Request) {
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
        if reqHost + reqUri == ``{
            fmt.Fprint(w,`empty url!`)
            return
        }
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
                tmpFile := s.path + `/tmp/` + fileHeader.Filename
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
                tmpPath := s.path + `/tmp`
                os.RemoveAll(tmpPath)
                os.Mkdir(tmpPath,0777)
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

func (s *server)clientInfo(w http.ResponseWriter,r *http.Request)  {
    clientInfo := make(map[string]interface{})
    clientInfo[`id`] = s.config.ClientID
    clientInfo[`secret`] = s.config.ClientSecret
    fmt.Fprint(w, response(clientInfo))
}

// 上传数据库comment
func (s *server)importDBComments(w http.ResponseWriter,r *http.Request){
    jsonInput,_ := ioutil.ReadAll(r.Body)
    var DB  dbConfig
    json.Unmarshal(jsonInput,&DB)
    DBStatement := fmt.Sprintf(`%s:%s@tcp(%s:%d)/information_schema?charset=utf8`,DB.Username,DB.Password,DB.Host,DB.Port)
    db,err := sql.Open(`mysql`,DBStatement)
    if err != nil {
        fmt.Println(err)
        fmt.Fprint(w,errResponse(1,`连接本地数据库失败`))
        return
    }
    rows,err := db.Query(`SELECT COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT FROM COLUMNS WHERE COLUMN_COMMENT !="" AND TABLE_SCHEMA IN (`+ DB.Databases +`) ORDER BY COLUMN_COMMENT DESC`)
    if err != nil {
        log.Print(err,`exec failed`)
        return
    }
    collection := make(map[string]map[string]string)

    for rows.Next() {
        var COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT string
        rows.Scan(&COLUMN_NAME,&DATA_TYPE,&COLUMN_COMMENT)
        if err != nil {
            continue
        }
        DATA_TYPE = getType(DATA_TYPE)
        key := md5Helper(COLUMN_NAME + DATA_TYPE + COLUMN_COMMENT)
        item := make(map[string]string)
        item[`key`] = COLUMN_NAME
        item[`statement`] = COLUMN_COMMENT
        item[`type`] = DATA_TYPE
        item[`hash`] = key
        collection[string(key)] = item
    }
    l := len(collection)
    list := make([]map[string]string,l)
    i := 0
    for _,v := range collection {
        list[i] = v
       i++
    }
    fmt.Fprint(w,response(list))
}

func errResponse(code int,message string)(data string){
    res := make(map[string]interface{})
    res[`code`] = code
    res[`message`] = message
    b,_ := json.Marshal(res)
    data = string(b)
    return
}

func response(data interface{})(result string){
    res := make(map[string]interface{})
    res[`code`] = 0
    res[`data`] = data
    resByte,_ := json.Marshal(res)
    result = string(resByte)
    return
}


func md5Helper (source string) (result string){
    Md5Inst := md5.New()
    Md5Inst.Write([]byte(source))
    res := Md5Inst.Sum([]byte(""))
    result = fmt.Sprintf("%x", res)
    return
}

func getType(DBFieldType string)(fieldType string){
    switch DBFieldType {
    case `int`,`bigint`,`tinyint`,`smallint`: fieldType = `int`
    case `varchar`,`text`,`longtext`,`char`:fieldType = `string`
    case `decimal`,`double`,`float`:fieldType = `float`
    default:fieldType = `string`
    }
    return
}


