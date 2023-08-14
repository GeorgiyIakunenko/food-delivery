import { protectedApi } from "@/api/base";

export const createOrder = async (order) => {
  try {
    const res = await protectedApi.post("/order/create", order);

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

export const getOrders = async () => {
  try {
    const response = await protectedApi.get("/order");

    if (response.status !== 200) {
      return {
        success: false,
        data: response.data,
      };
    }

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
