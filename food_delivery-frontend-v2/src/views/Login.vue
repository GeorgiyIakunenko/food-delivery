<script setup>
import Input from "@/components/Input.vue";
import Button from "@/components/Button.vue";
import {useLoginFormStore} from "@/store/loginForm";

const loginFormStore = useLoginFormStore();

const submitForm = async () => {
  console.log(loginFormStore.isLoginFormValid)
  if(loginFormStore.isLoginFormValid) {
    alert(loginFormStore.loginForm.email + " " + loginFormStore.loginForm.password);
  } else {
    alert("Form is invalid");
  }
}

</script>

<template>
  <main>
    <div class="container font-sans">
      <h1 class="text-neutral-800 font-bold mb-5 text-2xl text-center">Welcome Back</h1>
      <p class="text-neutral-100 font-normal mb-10 w-48 text-center mx-auto">Hello, sign in to continue!
        Or <router-link to="/register"><span class="text-primary-400">Create new account</span></router-link> </p>
      <div class="form w-4/5 mx-auto flex flex-col gap-7">
        <Input v-model="loginFormStore.loginForm.email"  label="Email Address" type="text" name="email"><span v-for="error in loginFormStore.loginFormValidation$.email.$errors" :key="error.$uid">{{error.$message}}</span></Input>
        <Input v-model="loginFormStore.loginForm.password"  label="Password" type="password" name="password"><span v-for="error in loginFormStore.loginFormValidation$.password.$errors" :key="error.$uid">{{error.$message}}</span></Input>
        <Button @click="submitForm" :disabled="!loginFormStore.isLoginFormValid" type="primary">Login</Button>
      </div>
      <router-link to="/reset-password" class="text-primary-400 text-sm text-center block mt-5">Forgot Password?</router-link>
    </div>
  </main>
</template>

<style scoped>

</style>