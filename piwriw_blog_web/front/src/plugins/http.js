import Vue from 'vue'
import axios from 'axios'

// axios请求地址
 axios.defaults.baseURL = 'http://localhost:8000'

Vue.prototype.$http = axios
