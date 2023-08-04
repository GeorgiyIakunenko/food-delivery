import { defineStore } from "pinia";
import { ref } from "vue";

export const useUserStore = defineStore("userStore", {
  state: () => {
    const UserID = ref(0);

    const accessToken = ref("");
    const refreshToken = ref("");

    return {
      accessToken,
      refreshToken,
      UserID,
    };
  },
  persist: true,
});
