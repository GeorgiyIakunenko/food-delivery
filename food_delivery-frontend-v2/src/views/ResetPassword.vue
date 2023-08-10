<script setup>
import Input from "@/components/Input.vue";
import Button from "@/components/Button.vue";

import { useResetFormStore } from "@/store/resetForm";
import { ref } from "vue";
import { resetPassword, resetPasswordRequest } from "@/api/user";
import Modal from "@/components/Modal.vue";
import router from "@/router/router";

let modalTitle = "success";
let modalMessage = "We send the code to your email address ";
let modalType = "success";
let modalOpen = ref(false);

const closeModal = () => {
  modalOpen.value = false;

  if (modalMessage === "Your password has been reset") {
    router.push("/login");
  }
};

const resetFormStore = useResetFormStore();

const isRequestForCodeSent = ref(false);

const submitResetRequestForm = async () => {
  console.log(resetFormStore.isResetEmailValid);
  if (resetFormStore.isResetEmailValid) {
    const res = await resetPasswordRequest(resetFormStore.resetForm.email);
    if (res.success === true) {
      modalTitle = "success";
      modalMessage =
        "We send the code to your email address " +
        "\n" +
        "It is valid for 10 minutes";
      modalType = "success";
      isRequestForCodeSent.value = true;
      modalOpen.value = true;
    } else {
      modalTitle = "Error";
      modalMessage = "Failed to send reset code";
      modalType = "error";
      modalOpen.value = true;
    }
  }
};
const submitForm = async () => {
  console.log(resetFormStore.isResetFormValid);
  if (resetFormStore.isResetFormValid) {
    const res = await resetPassword(
      resetFormStore.resetForm.email,
      resetFormStore.resetForm.code,
      resetFormStore.resetForm.password,
    );

    if (res.success === true) {
      modalTitle = "success";
      modalMessage = "Your password has been reset";
      modalType = "success";
      modalOpen.value = true;
    } else {
      modalTitle = "Error";
      modalMessage = "Code is invalid";
      modalType = "error";
      modalOpen.value = true;
    }
  } else {
  }
};
</script>

<template>
  <main>
    <div class="container font-sans">
      <h1 class="mb-5 text-center text-2xl font-bold text-neutral-800">
        Reset your password
      </h1>
      <p
        v-if="!isRequestForCodeSent"
        class="mx-auto mb-10 w-56 text-center font-normal text-neutral-100"
      >
        Write your email address and we will send you a code to reset your
        password
      </p>
      <div class="form mx-auto mt-10 flex w-4/5 flex-col gap-7 md:w-2/5">
        <Input
          v-if="!isRequestForCodeSent"
          v-model="resetFormStore.resetForm.email"
          label="Email Address"
          type="text"
          name="email"
          ><span
            v-for="error in resetFormStore.resetFormValidation$.email.$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Input
          v-if="isRequestForCodeSent"
          v-model="resetFormStore.resetForm.code"
          label="Code"
          type="text"
          name="code"
          ><span
            v-for="error in resetFormStore.resetFormValidation$.code.$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Input
          v-if="isRequestForCodeSent"
          v-model="resetFormStore.resetForm.password"
          label="Password"
          type="password"
          name="reset-password"
          ><span
            v-for="error in resetFormStore.resetFormValidation$.password
              .$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Input
          v-if="isRequestForCodeSent"
          v-model="resetFormStore.resetForm.confirmPassword"
          label="Confirm Password"
          type="password"
          name="confirm password"
          ><span
            v-for="error in resetFormStore.resetFormValidation$.confirmPassword
              .$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Button
          v-if="!isRequestForCodeSent"
          @click="submitResetRequestForm"
          :disabled="!resetFormStore.isResetEmailValid"
          type="primary"
          >Continue</Button
        >
        <Button
          v-else
          @click="submitForm"
          :disabled="!resetFormStore.isResetFormValid"
          type="primary"
          >Reset password</Button
        >
      </div>
      <router-link
        @click="isRequestForCodeSent = !isRequestForCodeSent"
        :to="isRequestForCodeSent ? '/reset-password' : '/login'"
        class="mt-5 block text-center text-sm"
        >Back</router-link
      >
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
