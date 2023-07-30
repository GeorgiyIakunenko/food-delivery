import { defineStore } from "pinia";
import {ref} from "vue";

export const useUserStore = defineStore("'userStore", () => {
   const UserID = ref(0);


    return {
        UserID,
    }
});