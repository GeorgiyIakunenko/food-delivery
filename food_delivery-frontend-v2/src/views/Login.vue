<script setup>
import Input from "@/components/Input.vue";
import Button from "@/components/Button.vue";
import { useLoginFormStore } from "@/store/loginForm";
import { login } from "@/api/api";
import Modal from "@/components/Modal.vue";
import { ref } from "vue";

const loginFormStore = useLoginFormStore();

let modalTitle = "Failed to login";
let modalMessage = "Email or password is not correct";
let modalType = "error";
let modalOpen = ref(false);

const submitForm = async () => {
  const res = await login(
    loginFormStore.loginForm.email,
    loginFormStore.loginForm.password,
  );
  if (res.success === true) {
    console.log(res.data);
  } else {
    modalOpen.value = true;
    alert(res.data);
  }
};
</script>

<template>
  <main>
    <div class="container font-sans">
      <h1 class="mb-5 text-center text-2xl font-bold text-neutral-800">
        Welcome Back
      </h1>
      <p class="mx-auto mb-10 w-48 text-center font-normal text-neutral-100">
        Hello, sign in to continue! Or
        <router-link to="/register"
          ><span class="text-primary-400">Create new account</span></router-link
        >
      </p>
      <div class="form mx-auto flex w-4/5 flex-col gap-7 md:w-2/5">
        <Input
          v-model="loginFormStore.loginForm.email"
          label="Email Address"
          type="text"
          name="email"
          ><span
            v-for="error in loginFormStore.loginFormValidation$.email.$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Input
          v-model="loginFormStore.loginForm.password"
          label="Password"
          type="password"
          name="password"
          ><span
            v-for="error in loginFormStore.loginFormValidation$.password
              .$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Button
          @click="submitForm"
          :disabled="!loginFormStore.isLoginFormValid"
          type="primary"
          >Login</Button
        >
      </div>
      <router-link
        to="/reset-password"
        class="mt-5 block text-center text-sm text-primary-400"
        >Forgot Password?</router-link
      >
    </div>
    <Modal
      :type="modalType"
      :title="modalTitle"
      @modalClose="modalOpen = false"
      v-bind:open="modalOpen"
      >{{ modalMessage }}</Modal
    >
  </main>
</template>

<style scoped></style>
