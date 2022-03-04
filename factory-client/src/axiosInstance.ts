import axios, { AxiosInstance, AxiosRequestConfig } from "axios";

const config: AxiosRequestConfig = {
  baseURL: "http://localhost:8080",
};
const instance: AxiosInstance = axios.create(config);

export default instance;
