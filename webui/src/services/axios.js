import axios from "axios";
import utils from "./utils"

const instance = axios.create({
  baseURL: __API_URL__,
  timeout: 1000 * 5
});

const setAuth = () => {
  const token = utils.getCurrentId();
  if (token) {
    // Remove 'Bearer' prefix, send just the token
    instance.defaults.headers.common['Authorization'] = token;
  } else {
    delete instance.defaults.headers.common['Authorization'];
  }
}

instance.interceptors.response.use(
  response => response,
  error => {
    if (error.response?.status === 401) {
      // Redirect to login on unauthorized
      window.location.href = '/#/login';
    }
    return Promise.reject(error);
  }
);

export {
  setAuth,
  instance as axios
}