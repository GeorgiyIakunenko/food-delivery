<script setup>
import Button from "@/components/Button.vue";
import Input from "@/components/Input.vue";
import { useUserStore } from "@/store/user";
import { ref } from "vue";

const userStore = useUserStore();
const isChangeMood = ref(false);
</script>

<template>
  <main>
    <div class="container">
      <div class="font-sans">
        <div class="mb-10 text-center text-2xl font-bold">Change password</div>
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
          <div class="flex flex-col gap-4">
            <Button class="w-full" type="primary" @click="">Change</Button>
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
  </main>
</template>

<style scoped></style>
