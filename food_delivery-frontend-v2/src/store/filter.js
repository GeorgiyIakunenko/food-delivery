import { defineStore } from "pinia";

export const useFilterStore = defineStore("filterStore", {
  state: () => {
    const filter = {
      search: "",
      categories: [],
      minPrice: 500,
      maxPrice: 2000,
      sortBy: "price",
      sortDirection: "",
      openNow: false,
    };

    return {
      filter,
    };
  },
  persist: true,
});
