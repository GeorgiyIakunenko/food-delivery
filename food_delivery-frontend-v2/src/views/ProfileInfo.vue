<script setup>
import { useUserStore } from "@/store/user";
import Input from "@/components/Input.vue";
import Button from "@/components/Button.vue";
import { ref } from "vue";
import router from "@/router/router";

const userStore = useUserStore();

const isChangeMood = ref(false);
</script>
<template>
  <main>
    <div class="container">
      <div class="font-sans">
        <div v-if="!isChangeMood" class="mb-10 text-center text-2xl font-bold">
          Profile Information
        </div>
        <div v-else class="mb-10 text-center text-2xl font-bold text-amber-700">
          Change profile information
        </div>
        <div class="mx-auto flex w-4/5 flex-col gap-7 md:w-3/5">
          <Input
            label="First Name"
            name="First Name"
            v-model="userStore.user.FirstName"
            :disabled="!isChangeMood"
          ></Input>
          <Input
            label="Last Name"
            name="Last Name"
            v-model="userStore.user.LastName"
            :disabled="!isChangeMood"
          ></Input>
          <Input
            label="Age"
            name="Age"
            v-model="userStore.user.Age"
            :disabled="!isChangeMood"
          ></Input>
          <Input
            label="Address"
            name="Address"
            v-model="userStore.user.Address"
            :disabled="!isChangeMood"
          ></Input>
          <div v-if="!isChangeMood" class="flex flex-col gap-4">
            <Button
              class="w-full"
              type="primary"
              @click="isChangeMood = !isChangeMood"
              >Edit</Button
            >
            <Button
              class="w-full"
              type="secondary"
              @click="$router.push('/profile')"
              >Back</Button
            >
          </div>
          <div v-else class="flex gap-3">
            <Button
              class="w-full"
              type="secondary"
              @click="isChangeMood = !isChangeMood"
              >Cancel</Button
            >
            <Button
              class="w-full"
              type="primary"
              @click="isChangeMood = !isChangeMood"
              >Save</Button
            >
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped></style>
