import { defineStore } from "pinia";

export const useFilterStore = defineStore("filterStore", {
  state: () => {
    const filter = {
      search: "",
      categories: [],
      minPrice: +0,
      maxPrice: +5000,
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
