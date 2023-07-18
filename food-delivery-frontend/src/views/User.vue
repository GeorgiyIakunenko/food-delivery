<script setup>
  import { RouterLink, RouterView } from 'vue-router'
  import {useUserStore} from "../stores/user";
  import Header from "../App.vue";

  const userStore = useUserStore()
</script>

<template>
  <main>
      <h1>user count {{ userStore.userCount}}</h1>
      <button @click="userStore.IncrementUserCount">increment</button>
  </main>
</template>

<style scoped>

</style>