import { defineStore } from "pinia";

export const useFilterStore = defineStore("filterStore", {
  state: () => {
    const filter = {
      search: "",
      categories: [],
      minPrice: 500,
      maxPrice: 2000,
      sortBy: "",
      sortDirection: "",
    };

    const setSearch = (search) => {
      filter.search = search;
    };

    const setCategories = (categories) => {
      filter.categories = categories;
    };

    const setPrice = (minPrice, maxPrice) => {
      filter.minPrice = minPrice;
      filter.maxPrice = maxPrice;
    };

    const setSort = (sortBy, sortDirection) => {
      filter.sortBy = sortBy;
      filter.sortDirection = sortDirection;
    };

    return {
      filter,
      setSearch,
      setCategories,
      setPrice,
      setSort,
    };
  },
  persist: true,
});
