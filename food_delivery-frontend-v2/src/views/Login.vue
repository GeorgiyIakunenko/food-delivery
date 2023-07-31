<script setup>
import Input from "@/components/Input.vue";
import {ref} from "vue";
import Button from "@/components/Button.vue";

const loginForm = ref({
  email: "goshan@gmail.com",
  password: "some password"
});



</script>

<template>
  <main>
    <div class="container font-sans">
      <h1 class="text-neutral-800 font-bold mb-5 text-2xl text-center">Welcome Back</h1>
      <p class="text-neutral-100 font-normal mb-7 w-52 text-center mx-auto">Hello, sign in to continue!
        Or <router-link to="/register"><span class="text-primary-400">Create new account</span></router-link> </p>
      <div class="form w-4/5 mx-auto flex flex-col gap-7">
        <Input v-model="loginForm.email"  label="Email Address" type="text" name="email" ></Input>
        <Input v-model="loginForm.password"  label="Password" type="password" name="password" ></Input>
        <Button :disabled="true" type="primary">Login</Button>
      </div>

    </div>
  </main>
</template>

<style scoped>

</style>