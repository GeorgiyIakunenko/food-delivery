<script setup>
import Button from "@/components/Button.vue";
</script>

<template>
  <main class="flex flex-col items-center justify-center p-24">
    <div class="flex flex-col items-center">
      <div class="font-sans text-xl font-medium text-red-500">
        <img class="mb-4" src="@/assets/images/404.svg" alt="page not found" />
        Page Not Found
      </div>
      <router-link to="/" exact
        ><Button class="mt-5" type="primary"
          >Go back to home</Button
        ></router-link
      >
    </div>
  </main>
</template>

<style scoped></style>
