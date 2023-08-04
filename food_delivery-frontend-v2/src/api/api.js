import axios from "axios";
import { useUserStore } from "@/store/user";

const api = axios.create({
  baseURL: "http://localhost:8080",
});

const protectedApi = axios.create({
  baseURL: "http://localhost:8080",
});

protectedApi.interceptors.request.use(
  (config) => {
    const token = useUserStore().accessToken;
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);

export const login = async (email, password) => {
  try {
    const response = await api.post("auth/login", {
      email: email,
      password: password,
    });

    // set the access token and refresh token

    ({
      access_token: useUserStore().accessToken,
      refresh_token: useUserStore().refreshToken,
    } = response.data);

    console.log(useUserStore().accessToken);

    return {
      success: true,
      data: response.data,
    };
  } catch (error) {
    return {
      success: false,
      data: error.response.statusText,
    };
  }
};

export const logout = async () => {
  try {
    const response = await protectedApi.post("auth/logout");

    // clear the access token and refresh token

    useUserStore().accessToken = "";
    useUserStore().refreshToken = "";

    return {
      success: true,
      data: response.data,
    };
  } catch (error) {
    return {
      success: false,
      data: error.response.statusText,
    };
  }
};
