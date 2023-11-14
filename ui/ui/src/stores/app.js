import { useStorage } from '@vueuse/core'



export  const  state = useStorage(
    'vblog',
    {
        is_login:false,
        refresh_token:'',
        access_token:'',
        role:0
    },
    localStorage,
    {mergeDefaults:true}
)