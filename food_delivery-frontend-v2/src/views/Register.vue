<script setup>
import Input from "@/components/Input.vue";
import Button from "@/components/Button.vue";
import { useRegisterFormsStore } from "@/store/registerForm";

const registerFormStore = useRegisterFormsStore();

const submitForm = async () => {
  if (registerFormStore.isRegisterFormValid) {
    alert(
      registerFormStore.registerForm.email +
        " " +
        registerFormStore.registerForm.password,
    );
  } else {
    alert("Form is invalid");
  }
};
</script>

<template>
  <main>
    <div class="container font-sans">
      <h1 class="mb-5 text-center text-2xl font-bold text-neutral-800">
        Hello! Create Account
      </h1>
      <p class="mx-auto mb-10 w-48 text-center font-normal text-neutral-100">
        Already have an account? Or
        <router-link to="/login"
          ><span class="text-primary-400"> Sign in</span></router-link
        >
      </p>
      <div class="form mx-auto flex w-4/5 flex-col gap-4 md:w-2/5">
        <Input
          v-model="registerFormStore.registerForm.email"
          label="Email Address"
          type="text"
          name="email"
          ><span
            v-for="error in registerFormStore.registerFormValidation$.email
              .$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Input
          v-model="registerFormStore.registerForm.firstName"
          label="First Name"
          type="text"
          name="firstName"
          ><span
            v-for="error in registerFormStore.registerFormValidation$.firstName
              .$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Input
          v-model="registerFormStore.registerForm.lastName"
          label="Last Name"
          type="text"
          name="lastName"
          ><span
            v-for="error in registerFormStore.registerFormValidation$.lastName
              .$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Input
          v-model="registerFormStore.registerForm.username"
          label="Username"
          type="text"
          name="username"
          ><span
            v-for="error in registerFormStore.registerFormValidation$.username
              .$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Input
          v-model="registerFormStore.registerForm.age"
          label="Age"
          type="number"
          name="age"
          ><span
            v-for="error in registerFormStore.registerFormValidation$.age
              .$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Input
          v-model="registerFormStore.registerForm.address"
          label="Address"
          type="text"
          name="address"
          ><span
            v-for="error in registerFormStore.registerFormValidation$.address
              .$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Input
          v-model="registerFormStore.registerForm.password"
          label="Password"
          type="password"
          name="password"
          ><span
            v-for="error in registerFormStore.registerFormValidation$.password
              .$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Input
          v-model="registerFormStore.registerForm.confirmPassword"
          label="Confirm Password"
          type="password"
          name="confirmPassword"
          ><span
            v-for="error in registerFormStore.registerFormValidation$
              .confirmPassword.$errors"
            :key="error.$uid"
            >{{ error.$message }}</span
          ></Input
        >
        <Button
          @click="submitForm"
          :disabled="!registerFormStore.isRegisterFormValid"
          type="primary"
          class="mt-10"
          >Register</Button
        >
      </div>
    </div>
  </main>
</template>

<style scoped></style>
