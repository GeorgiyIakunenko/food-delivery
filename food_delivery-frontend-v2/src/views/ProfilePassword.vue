<script setup>
import Button from "@/components/Button.vue";
import Input from "@/components/Input.vue";
import { useUserStore } from "@/store/user";
import { computed, reactive, ref, watch } from "vue";
import useValidate from "@vuelidate/core";
import { helpers, minLength, required, sameAs } from "@vuelidate/validators";
import { changePassword } from "@/api/user";
import Modal from "@/components/Modal.vue";
import router from "@/router/router";

let modalTitle = "success";
let modalMessage = "We send the code to your email address ";
let modalType = "success";
let modalOpen = ref(false);

const closeModal = () => {
  modalOpen.value = false;

  if (modalMessage === "success") {
    router.push("/profile");
  }
};

const changePasswordForm = reactive({
  currentPassword: "",
  newPassword: "123456",
  confirmPassword: "123456",
});

const changePasswordFormRules = computed(() => {
  return {
    currentPassword: {
      required: helpers.withMessage("This line can't be empty", required),
    },
    newPassword: {
      required: helpers.withMessage("This line can't be empty", required),
      minLength: helpers.withMessage(
        "Password should contain at least 6 characters",
        minLength(6),
      ),
    },
    confirmPassword: {
      sameAs: helpers.withMessage(
        "The passwords should be equal",
        sameAs(changePasswordForm.newPassword),
      ),
    },
  };
});

let isChangePasswordFormValid = ref(false);

const changePasswordFormValidation$ = useValidate(
  changePasswordFormRules,
  changePasswordForm,
);

watch(changePasswordForm, async () => {
  isChangePasswordFormValid.value =
    await changePasswordFormValidation$.value.$validate();
});

const submitForm = () => {
  if (isChangePasswordFormValid.value) {
    const res = changePassword(
      changePasswordForm.currentPassword,
      changePasswordForm.newPassword,
    );
    if (res.success === true) {
      modalTitle = "success";
      modalMessage = "Your password has been changed";
      modalType = "success";
      modalOpen.value = true;
    } else {
      modalTitle = "Error";
      modalMessage = "Failed to change password";
      modalType = "error";
      modalOpen.value = true;
    }
  }
};
</script>

<template>
  <main>
    <div class="container">
      <div class="font-sans">
        <div class="mb-10 text-center text-2xl font-bold">Change password</div>
        <div class="mx-auto flex w-4/5 flex-col gap-7 md:w-3/5">
          <Input
            label="Current password"
            name="Current password"
            v-model="changePasswordForm.currentPassword"
            ><span
              v-for="error in changePasswordFormValidation$.currentPassword
                .$errors"
              :key="error.$uid"
              >{{ error.$message }}</span
            ></Input
          >
          <Input
            label="New password"
            name="New password"
            v-model="changePasswordForm.newPassword"
            ><span
              v-for="error in changePasswordFormValidation$.newPassword.$errors"
              :key="error.$uid"
              >{{ error.$message }}</span
            ></Input
          >
          <Input
            label="Confirm new password"
            name="Confirm new password"
            v-model="changePasswordForm.confirmPassword"
          >
            <span
              v-for="error in changePasswordFormValidation$.confirmPassword
                .$errors"
              :key="error.$uid"
              >{{ error.$message }}</span
            >
          </Input>
          <div class="flex flex-col gap-4">
            <Button
              class="w-full"
              type="primary"
              :disabled="!isChangePasswordFormValid"
              @click="submitForm"
              >Change</Button
            >
            <Button
              class="w-full"
              type="secondary"
              @click="$router.push('/profile')"
              >Back</Button
            >
          </div>
        </div>
      </div>
    </div>
    <Modal
      :type="modalType"
      :title="modalTitle"
      @modalClose="closeModal"
      v-bind:open="modalOpen"
      >{{ modalMessage }}</Modal
    >
  </main>
</template>

<style scoped></style>
