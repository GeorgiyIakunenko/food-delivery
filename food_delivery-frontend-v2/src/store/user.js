import { defineStore } from "pinia";
import { reactive, ref } from "vue";

export const useUserStore = defineStore("userStore", {
  state: () => {
    const accessToken = ref("");
    const refreshToken = ref("");

    const user = reactive({
      id: 0,
      FirstName: "",
      LastName: "",
      Email: "",
      Age: 0,
      Phone: "",
      Address: "",
      Username: "",
    });

    const setUser = (data) => {
      user.id = data.id;
      user.FirstName = data.firstName;
      user.LastName = data.lastName;
      user.Email = data.email;
      user.Age = data.age;
      user.Phone = data.phone;
      user.Address = data.address;
      user.Username = data.username;
    };

    return {
      accessToken,
      refreshToken,
      user,
      setUser,
    };
  },
  persist: true,
});
