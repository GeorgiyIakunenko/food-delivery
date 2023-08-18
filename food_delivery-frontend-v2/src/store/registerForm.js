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

export const useRegisterFormsStore = defineStore("registerFormsStore", {
  state: () => {
    // register form

    const registerForm = reactive({
      firstName: "",
      lastName: "",
      email: "",
      username: "",
      phone: "",
      age: "",
      address: "",
      password: "",
      confirmPassword: "",
    });

    const registerFormRules = computed(() => {
      return {
        email: {
          required: helpers.withMessage("Email is required", required),
          email: helpers.withMessage("Email is invalid", email),
        },
        firstName: {
          required: helpers.withMessage("First name is required", required),
        },
        lastName: {
          required: helpers.withMessage("Last name is required", required),
        },
        username: {
          required: helpers.withMessage("Username is required", required),
        },
        phone: {
          required: helpers.withMessage("Phone is required", required),
        },
        age: {
          required: helpers.withMessage("Age is required", required),
        },
        address: {
          required: helpers.withMessage("Address is required", required),
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
            sameAs(registerForm.password),
          ),
        },
      };
    });

    const registerFormValidation$ = useValidate(
      registerFormRules,
      registerForm,
    );

    let isRegisterFormValid = ref(false);

    watch(registerForm, async () => {
      isRegisterFormValid.value =
        await registerFormValidation$.value.$validate();
    });

    return {
      registerForm,
      registerFormValidation$,
      isRegisterFormValid,
    };
  },
});
