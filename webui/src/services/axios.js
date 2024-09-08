import axios from "axios";
import utils from "./utils"

const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

const setAuth = () => {
    instance.defaults.headers.common['Authorization'] = 'Bearer ' + utils.getCurrentId()
}

export {
	setAuth,
	instance as axios
}
