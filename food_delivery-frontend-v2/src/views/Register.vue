<script setup>
import Input from "@/components/Input.vue";
import Button from "@/components/Button.vue";
import {useFormsStore} from "@/store/forms";


const formsStore = useFormsStore();

const submitForm = async () => {
  const isValid = await formsStore.registerFormValidation$.value.$validate();
  if (isValid) {
    alert(formsStore.registerForm.email + " " + formsStore.registerForm.password);
  } else {
    alert("Form is invalid");
  }
}
</script>

<template>
  <main>
    <div class="container font-sans">
      <h1 class="text-neutral-800 font-bold mb-5 text-2xl text-center">Hello! Create Account</h1>
      <p class="text-neutral-100 font-normal mb-10 w-48 text-center mx-auto">Already have an account?
        Or
        <router-link to="/login"><span class="text-primary-400"> Sign in</span></router-link>
      </p>
      <div class="form w-4/5 mx-auto flex flex-col gap-4">
        <Input v-model="formsStore.registerForm.email" label="Email Address" type="text" name="email"><span
            v-for="error in formsStore.registerFormValidation$.email.$errors" :key="error.$uid">{{ error.$message }}</span></Input>
        <Input v-model="formsStore.registerForm.firstName" label="First Name" type="text" name="firstName"><span
            v-for="error in formsStore.registerFormValidation$.firstName.$errors" :key="error.$uid">{{ error.$message }}</span></Input>
        <Input v-model="formsStore.registerForm.lastName" label="Last Name" type="text" name="lastName"><span
            v-for="error in  formsStore.registerFormValidation$.lastName.$errors" :key="error.$uid">{{ error.$message }}</span></Input>
        <Input v-model="formsStore.registerForm.username" label="Username" type="text" name="username"><span
            v-for="error in formsStore.registerFormValidation$.username.$errors" :key="error.$uid">{{ error.$message }}</span></Input>
        <Input v-model="formsStore.registerForm.age" label="Age" type="number" name="age"><span v-for="error in formsStore.registerFormValidation$.age.$errors"
                                                                                     :key="error.$uid">{{
            error.$message
          }}</span></Input>
        <Input v-model="formsStore.registerForm.address" label="Address" type="text" name="address"><span
            v-for="error in formsStore.registerFormValidation$.address.$errors" :key="error.$uid">{{ error.$message }}</span></Input>
        <Input v-model="formsStore.registerForm.password" label="Password" type="password" name="password"><span
            v-for="error in formsStore.registerFormValidation$.password.$errors" :key="error.$uid">{{ error.$message }}</span></Input>
        <Input v-model="formsStore.registerForm.confirmPassword" label="Confirm Password" type="password"
               name="confirmPassword"><span v-for="error in formsStore.registerFormValidation$.confirmPassword.$errors"
                                            :key="error.$uid">{{ error.$message }}</span></Input>
        <Button @click="submitForm" :disabled="!formsStore.isRegisterFormValid" type="primary">Register</Button>
      </div>
      <router-link to="/reset-password" class="text-primary-400 text-sm text-center block mt-5">Forgot Password?
      </router-link>
    </div>
  </main>
</template>

<style scoped>

</style>