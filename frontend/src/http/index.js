import axios from 'axios'
import router from '../router/index.js'

axios.defaults.timeout = 5000;
axios.defaults.baseURL = '';
axios.interceptors.request.use(
    config => {
let expire_time = localStorage.getItem('expire_time')
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
        console.log(error)
    switch (error.response) {
        case 401:
            // 401 清除token信息并跳转到登录页面
            router.replace({
                path: 'login',
                query: {redirect: router.currentRoute.fullPath}
            })
    }
    console.log(error.response)
    return  Promise.reject(error.response.data)
    }

)

export default axios;