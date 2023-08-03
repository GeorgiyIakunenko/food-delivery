import { defineStore } from "pinia";
import { computed, reactive, ref, watch } from "vue";
import { email, helpers, minLength, required } from "@vuelidate/validators";
import useValidate from "@vuelidate/core";

export const useLoginFormStore = defineStore("loginFormStore", {
  state: () => {
    // login form
    const loginForm = reactive({
      email: "goshan3097@gmail",
      password: "666666",
    });

    const loginFormRules = computed(() => {
      return {
        email: {
          required: helpers.withMessage("Email is required", required),
          email: helpers.withMessage("Email is invalid", email),
        },
        password: {
          required: helpers.withMessage("This line can't be empty", required),
          minLength: helpers.withMessage(
            "Password should contain at least 6 characters",
            minLength(6),
          ),
        },
      };
    });

    const loginFormValidation$ = useValidate(loginFormRules, loginForm);

    let isLoginFormValid = ref(false);

    watch(loginForm, async () => {
      isLoginFormValid.value = await loginFormValidation$.value.$validate();
    });

    return {
      loginForm,
      loginFormValidation$,
      isLoginFormValid,
    };
  },
});
