import { defineStore } from "pinia";
import { reactive } from "vue";

export const useFilterStore = defineStore("filterStore", {
  state: () => {
    const filter = reactive({
      search: "",
      categories: [],
      minPrice: +0,
      maxPrice: +5000,
      sortBy: "price",
      sortDirection: "",
      openNow: false,
    });

    const resetFilter = () => {
      filter.search = "";
      filter.categories = [];
      filter.minPrice = +0;
      filter.maxPrice = +5000;
      filter.sortBy = "price";
      filter.sortDirection = "";
      filter.openNow = false;
    };

    return {
      filter,
      resetFilter,
    };
  },
  persist: true,
});
