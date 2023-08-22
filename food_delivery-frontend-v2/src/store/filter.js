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
      resetCategoryFilter();
      resetPriceFilter();
      resetSortFilter();
      resetOpenNowFilter();
        resetSearchFilter();
    };

    const isCategoryFiltered = () => {
      return filter.categories.length > 0;
    };

    const resetCategoryFilter = () => {
      filter.categories = [];
    };

    const isPriceFiltered = () => {
      return filter.minPrice > 0 || filter.maxPrice < 5000;
    };

    const resetPriceFilter = () => {
      filter.minPrice = +0;
      filter.maxPrice = +5000;
    };

    const isSortFiltered = () => {
      return filter.sortBy !== "price" || filter.sortDirection !== "";
    };

    const resetSortFilter = () => {
      filter.sortBy = "price";
      filter.sortDirection = "";
    };

    const isOpenNowFiltered = () => {
      return filter.openNow;
    };

    const resetOpenNowFilter = () => {
      filter.openNow = false;
    };

    const isSearchFiltered = () => {
        return filter.search !== "";
    }

    const resetSearchFilter = () => {
        filter.search = "";
    }

    const isFiltered = () => {
      return (
        isSortFiltered() ||
        isCategoryFiltered() ||
        isPriceFiltered() ||
        isOpenNowFiltered() || isSearchFiltered()
      );
    };

    return {
      filter,
      resetFilter,
      isCategoryFiltered,
      isPriceFiltered,
      isSortFiltered,
      isOpenNowFiltered,
      isFiltered,
      resetCategoryFilter,
      resetPriceFilter,
      resetSortFilter,
      resetOpenNowFilter,
    };
  },
  persist: true,
});
