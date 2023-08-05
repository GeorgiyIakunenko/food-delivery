import axios from "axios";
import { useUserStore } from "@/store/user";
import router from "@/router/router";

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
      config.headers.Authorization = `bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);

protectedApi.interceptors.response.use(
  (response) => {
    console.log("interceptor response", response);
    return response;
  },
  async (error) => {
    const originalRequest = error.config;

    if (error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      try {
        const res = await refresh();
        console.log("interceptor response", res.data);
        if (res.data.status !== 200) {
          useUserStore().accessToken = "";
          useUserStore().refreshToken = "";
          await router.push("/login");
          return Promise.reject(error);
        }
        ({
          access_token: useUserStore().accessToken,
          refresh_token: useUserStore().refreshToken,
        } = res.data);

        originalRequest.headers.Authorization = `Bearer ${
          useUserStore().accessToken
        }`;
        return protectedApi(originalRequest);
      } catch (error) {
        return Promise.reject(error);
      }
    }
    return Promise.reject(error);
  },
);

export const refresh = async () => {
  try {
    console.log("refresh token", useUserStore().refreshToken);
    const response = await api.post(
      "auth/refresh",
      {},
      {
        headers: {
          Authorization: `bearer ${useUserStore().refreshToken}`,
        },
      },
    );
    // set the access token and refresh token
    ({
      access_token: useUserStore().accessToken,
      refresh_token: useUserStore().refreshToken,
    } = response.data);
    console.log("refresh response", response.data);
    return {
      success: true,
      data: response.data,
    };
  } catch (error) {
    console.log("error", error.response.data);
    return {
      success: false,
      data: error.response.data,
    };
  }
};

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

    return {
      success: true,
      data: response.data,
    };
  } catch (error) {
    return {
      success: false,
      data: error.response.data,
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
      data: error.response.data,
    };
  }
};
