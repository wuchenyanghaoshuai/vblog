import axios        from "axios";
import {Message} from "@arco-design/web-vue";


var instance = axios.create({
    baseURL: '',
    timeout: 5000,
    headers: {'Content-Type': 'application/json'}
})

export default instance


instance.interceptors.response.use(
    (resp)=>{
       // console.log(resp)
        return resp
    },
    (error)=>{
        let msg = error.message

        if (error.response.data && error.response.data.message){
            msg = error.response.data.message
            switch (error.response.data.code){
                //token过期，重新登陆
                case 5001:
                    window.location.assign('/login')

                    break
                default:
                    break
            }
            return Promise.reject(error.response.data)
        }
        Message.error(`${msg}`)

    }
)