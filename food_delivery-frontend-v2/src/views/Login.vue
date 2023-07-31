<script setup>
import Input from "@/components/Input.vue";
import {ref} from "vue";

const loginForm = ref({
  email: "",
  password: ""
});

</script>

<template>
  <main>
    <div class="container">
      Login
      <div class="form">
        <Input v-model:value="loginForm.email"  label="Email Address" type="text" value="email" ></Input>
        <Input v-model:value="loginForm.password"  label="Password" type="password" value="password" ></Input>
      </div>

    </div>
  </main>
</template>

<style scoped>

</style>