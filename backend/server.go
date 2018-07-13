package main

import(
  "net/http"
    "net/url"
    "net/http/httputil"
    "fmt"
)

func main() {

mux := http.NewServeMux()

hFile := http.FileServer(http.Dir(`/Users/zhu/project/docMakerClient/frontend/dist`))

 mux.HandleFunc(`/api/`,apiProxy)
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
