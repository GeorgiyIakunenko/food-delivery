import { api } from "@/api/base";

export const getAllProducts = async () => {
  try {
    const res = await api.get("products");

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

export const getProductsByCategoryAndSupplierIDs = async (
  categoryId,
  supplierId,
) => {
  try {
    const res = await api.get(
      `categories/${categoryId}/suppliers/${supplierId}/products`,
    );

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
