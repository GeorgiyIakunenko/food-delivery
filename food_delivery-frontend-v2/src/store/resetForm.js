import { defineStore } from "pinia";
import { computed, reactive, ref, watch } from "vue";
import {
  email,
  helpers,
  minLength,
  required,
  sameAs,
} from "@vuelidate/validators";
import useValidate from "@vuelidate/core";

export const useResetFormStore = defineStore("resetFormStore", {
  state: () => {
    // login form
    const resetForm = reactive({
      email: "",
      code: "",
      password: "",
      confirmPassword: "",
    });

    const resetFormRules = computed(() => {
      return {
        email: {
          required: helpers.withMessage("Email is required", required),
          email: helpers.withMessage("Email is invalid", email),
        },
        code: {
          required: helpers.withMessage("This line can't be empty", required),
        },
        password: {
          required: helpers.withMessage("This line can't be empty", required),
          minLength: helpers.withMessage(
            "Password should contain at least 6 characters",
            minLength(6),
          ),
        },
        confirmPassword: {
          sameAs: helpers.withMessage(
            "The passwords should be equal",
            sameAs(resetForm.password),
          ),
        },
      };
    });

    const resetFormValidation$ = useValidate(resetFormRules, resetForm);

    let isResetFormValid = ref(false);
    let isResetEmailValid = ref(false);

    watch(resetForm, async () => {
      isResetFormValid.value = await resetFormValidation$.value.$validate();
      isResetEmailValid.value =
        await resetFormValidation$.value.email.$validate();
    });

    return {
      resetForm,
      resetFormValidation$,
      isResetFormValid,
      isResetEmailValid,
    };
  },
  persist: true,
});
