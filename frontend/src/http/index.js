import axios from 'axios'
import router from '../router/index.js'

axios.defaults.timeout = 10000;
axios.defaults.baseURL = '';
axios.interceptors.request.use(
    config => {
var expire_time = localStorage.getItem('expire_time')
    if(expire_time){
        config.headers.Authorization = localStorage.getItem('token_type') + ' ' +localStorage.getItem('access_token')
}
return config
    }, error=>{
    return Promise.reject(error)
});

axios.interceptors.response.use(
    response=>{return response},
    error=>{
    switch (error.response.status) {
        case 401:
            // 如果页面里面有多个请求，会redirect两次
          if(router.currentRoute.name !=='Login' ){
            // 401 清除token信息并跳转到登录页面
            router.replace({
              path: '/login',
              query: {redirect: router.currentRoute.fullPath}
            })
          }

    }
    return  Promise.reject(error.response.data)
    }

)

export default axios;