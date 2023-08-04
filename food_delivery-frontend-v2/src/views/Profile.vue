<script setup>
import Button from "@/components/Button.vue";
import { logout } from "@/api/api";

const logoutUser = async () => {
  const result = await logout();
  if (result.success === true) {
    console.log(result.data);
  } else {
    alert(result.data);
  }
};
</script>

<template>
  <main>
    <div class="container">
      <div class="">Profile</div>
      <Button @click="logoutUser">Logout</Button>
    </div>
  </main>
</template>

<style scoped></style>
