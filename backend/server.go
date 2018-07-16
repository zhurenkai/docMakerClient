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

func main() {

mux := http.NewServeMux()

hFile := http.FileServer(http.Dir(`../frontend/dist`))

 mux.HandleFunc(`/api/`,apiProxy)
 mux.HandleFunc(`/client-api/`,requestProxy)
 mux.Handle(`/`,hFile)


  http.ListenAndServe(":9999",mux)

}

func apiProxy(w http.ResponseWriter,r *http.Request) {
    remote,err := url.Parse(`http://docmaker-server.randqun.com`)
    if err !=nil {
       fmt.Println(err)
    }

    proxy :=httputil.NewSingleHostReverseProxy(remote)
    // 重写rui
    b := []byte(r.RequestURI)
    r.RequestURI = string(b[4:])
    r.URL = &url.URL{Path:r.RequestURI}
    // 重设host否则就是反向代理，会将原来的请求带过去导致
    r.Host = `docmaker-server.randqun.com`
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
        fmt.Println(jsonInput)
        var jsonContainer map[string]string
        err := json.Unmarshal(jsonInput,&jsonContainer)
        fmt.Println(jsonContainer,err)
        reqHeaders = jsonContainer[`request_headers`]
        reqMethod = jsonContainer[`request_method`]
        reqHost = jsonContainer[`request_host`]
        reqUri = jsonContainer[`request_uri`]
        jsonData := jsonContainer[`json_input`]
        //var data interface{}
        //data = json.Unmarshal([]byte(jsondata),&data)
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
            for field := range r.MultipartForm.File{

                go func(ch chan string) {
                file, _, _   := r.FormFile(field)
                tmpFileName := fmt.Sprintf(`%d%d`, rand.Intn(999999999) ,time.Now().Unix())
                f,err := os.OpenFile(`tmp/` + tmpFileName,os.O_WRONLY|os.O_CREATE, 0666)
                if err !=nil {
                    fmt.Println(err)
                }
                io.Copy(f,file)
                f.Close()
                writer.CreateFormFile(field,`tmp/` +tmpFileName)
                ch<-`tmp/` +tmpFileName
                }(ch)
            }
            //var tmpFile  [tmpFileNum]string
           for i:=0;i<tmpFileNum;i++{
               // todo delete tmpFile
               fmt.Println(<-ch)
           }
        }
        writer.Close()
        contentType := writer.FormDataContentType()
        writer.Close()
        fmt.Println(data)
        req,err = http.NewRequest(reqMethod,reqHost + reqUri,data)
        req.Header.Add(`Content-Type`,contentType)
        req.Header.Set("Connection", "Keep-Alive")
    }


    // process header
    var headersContainer map[string]string
    json.Unmarshal([]byte(reqHeaders),&headersContainer)
    for k,v:= range headersContainer {
        req.Header.Add(k,v)
    }
    //fmt.Println(r.PostForm)

    resp,err :=client.Do(req)
    bodyBytes,_ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    fmt.Fprint(w,string(bodyBytes))
    if err != nil {
        fmt.Println(err)  
    }
    //fmt.Println(r.Form)
}
