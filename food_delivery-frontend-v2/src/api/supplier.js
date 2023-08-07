import { api } from "@/api/base";

export const getAllSuppliers = async () => {
  try {
    const res = await api.get("suppliers");

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
