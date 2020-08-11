import axios from "axios";

const httpService = axios.create({
    baseURL: "http://localhost:8080/",
    headers: {
        "Content-type": "application/json"
    }
});

// request interceptor
httpService.interceptors.request.use(config => {
    // ...
    return config
}, err => {
    return Promise.reject(err)
})

// response interceptor
httpService.interceptors.response.use(response => {
    // ...
    return response
}, err => {
    return Promise.reject(err)
})

export default httpService
