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

export const getFilteredProducts = async (filter) => {
  try {
    const queryParams = new URLSearchParams({
      search: filter.search,
      order_by: filter.sortBy,
      open_now: filter.openNow,
      sort_direction: filter.sortDirection,
      category_ids: filter.categories.join(","),
      min_price: filter.minPrice,
      max_price: filter.maxPrice,
    }).toString();
    const res = await api.get(`products/filter?${queryParams}`);
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
