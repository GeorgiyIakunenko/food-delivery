import { useUserStore } from "@/store/user";
import { api, protectedApi } from "@/api/base";

export const register = async (
  firstName,
  lastName,
  userName,
  age,
  email,
  phone,
  password,
  address,
) => {
  try {
    const res = await api.post("auth/register", {
      first_name: firstName,
      last_name: lastName,
      username: userName,
      age: +age,
      email: email,
      phone: phone,
      password: password,
      address: address,
    });

    if (res.status !== 200) {
      return {
        success: false,
        data: res.data,
      };
    }

    return {
      success: true,
      data: res.data,
    };
  } catch (error) {
    return {
      success: false,
      data: error.response.data,
    };
  }
};

export const resetPasswordRequest = async (email) => {
  try {
    const res = await api.post("auth/reset-password", {
      email: email,
    });

    if (res.status !== 200) {
      return {
        success: false,
        data: res.data,
      };
    }

    return {
      success: true,
      data: res.data,
    };
  } catch (error) {
    return {
      success: false,
      data: error.response.data,
    };
  }
};

export const resetPassword = async (email, code, password) => {
  try {
    const res = await api.post("auth/submit-code", {
      email: email,
      reset_code: code,
      new_password: password,
    });

    if (res.status !== 200) {
      return {
        success: false,
        data: res.data,
      };
    }

    return {
      success: true,
      data: res.data,
    };
  } catch (error) {
    return {
      success: false,
      data: error.response.data,
    };
  }
};

export const changePassword = async (oldPassword, newPassword) => {
  try {
    const res = await protectedApi.post("user/profile/password", {
      old_password: oldPassword,
      new_password: newPassword,
    });

    if (res.status !== 200) {
      return {
        success: false,
        data: res.data,
      };
    }

    return {
      success: true,
      data: res.data,
    };
  } catch (error) {
    return {
      success: false,
      data: error.response.data,
    };
  }
};

export const login = async (email, password) => {
  try {
    const res = await api.post("auth/login", {
      email: email,
      password: password,
    });

    // set the access token and refresh token

    ({
      access_token: useUserStore().accessToken,
      refresh_token: useUserStore().refreshToken,
    } = res.data);

    return {
      success: true,
      data: res.data,
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
    const res = await protectedApi.post("auth/logout");

    // clear the access token and refresh token

    useUserStore().accessToken = "";
    useUserStore().refreshToken = "";

    return {
      success: true,
      data: res.data,
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

    if (response.status !== 200) {
      return {
        success: false,
        data: response.data,
      };
    }

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
