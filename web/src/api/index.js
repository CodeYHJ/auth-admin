import axios from 'axios'

const request = axios.create({ baseURL: import.meta.env.API_URL })

request.interceptors.response.use(res => {
    const { data, code, msg } = res.data
    if (code === 200) {
        return data
    }
    return Promise.reject("msg")
})

export const getAccountList = () => {
    return request.get("/api/account/list")
}

export const editAccount = (data) => {
    return request.post("/api/account/edit", data)
}

// export const addAccountList = () => {
//     return request.post("/api/account/add")
// }


export const getClientList = () => {
    return request.get("/api/client/list")
}

export const getTokenList = () => {
    return request.get("/api/token/list")
}