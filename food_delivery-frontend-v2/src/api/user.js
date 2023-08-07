import { useUserStore } from "@/store/user";
import { api, protectedApi } from "@/api/base";

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

export const getProfile = async () => {
  try {
    const response = await protectedApi.get("user/profile");
    useUserStore().setUser(response.data);
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
