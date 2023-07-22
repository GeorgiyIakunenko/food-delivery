<script setup>
import Footer from "@/components/Footer.vue";
import Header from "@/components/Header.vue";
import {useUserStore} from "@/stores/user";

const userStore = useUserStore()
</script>

<template>
  <Header />
  <main>
    HomePage
  </main>
  <Footer />
</template>
