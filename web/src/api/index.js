import axios from 'axios'

const request = axios.create({ baseURL: "/api" })

request.interceptors.response.use(res => {
    const { data, code, msg } = res.data
    if (code === 200) {
        return data
    }
    return Promise.reject("msg")
})

export const getAccountList = () => {
    return request.get("/admin/account/list")
}

export const editAccount = (data) => {
    return request.post("/admin/account/edit", data)
}

// export const addAccountList = () => {
//     return request.post("/admin/account/add")
// }


export const getClientList = () => {
    return request.get("/admin/client/list")
}

export const getTokenList = () => {
    return request.get("/admin/token/list")
}