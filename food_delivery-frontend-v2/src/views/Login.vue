<script setup>
import Input from "@/components/Input.vue";
import {computed, reactive, ref, watch} from "vue";
import Button from "@/components/Button.vue";
import useValidate from "@vuelidate/core";
import {required, email, minLength} from "@vuelidate/validators";

const loginForm = reactive({
  email: "",
  password: ""
});

const loginFormRules = computed(() => {
  return {
    email: {required, email},
    password: {required, minLength: minLength(6)}
  }
})

const v$ = useValidate(loginFormRules, loginForm);

const submitForm = async () => {
  const isValid = await v$.value.$validate();
  if(isValid) {
    alert(loginForm.email + " " + loginForm.password);
  } else {
    alert("Form is invalid");
  }
}

let isFormValid = ref(false);

watch(loginForm, async () => {
  isFormValid =  await v$.value.$validate();
  // console.log(isFormValid);
})

</script>

<template>
  <main>
    <div class="container font-sans">
      <h1 class="text-neutral-800 font-bold mb-5 text-2xl text-center">Welcome Back</h1>
      <p class="text-neutral-100 font-normal mb-10 w-48 text-center mx-auto">Hello, sign in to continue!
        Or <router-link to="/register"><span class="text-primary-400">Create new account</span></router-link> </p>
      <div class="form w-4/5 mx-auto flex flex-col gap-7">
        <Input v-model="loginForm.email"  label="Email Address" type="text" name="email"><span v-for="error in v$.email.$errors" :key="error.$uid">{{error.$message}}</span></Input>
        <Input v-model="loginForm.password"  label="Password" type="password" name="password"><span v-for="error in v$.password.$errors" :key="error.$uid">{{error.$message}}</span></Input>
        <Button @click="submitForm" :disabled="!isFormValid" type="primary">Login</Button>
      </div>
      <router-link to="/reset-password" class="text-primary-400 text-sm text-center block mt-5">Forgot Password?</router-link>
    </div>
  </main>
</template>

<style scoped>

</style>