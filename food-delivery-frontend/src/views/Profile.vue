<script setup>
import Footer from "@/components/Footer.vue";
import Header from "@/components/Header.vue";
import {useUserStore} from "@/stores/user";
import {onMounted} from "vue";
import {getUserData} from "@/api/api";

const userStore = useUserStore()

onMounted(() => {
  getUserData()
})

</script>
<template>
  <Header />
  <main>
    Profile
    <div class="">{{userStore.user.id}}</div>
    <div class="">{{userStore.user.name}}</div>
    <div class="">{{userStore.user.email}}</div>
    <div class="">{{userStore.user.age}}</div>
    <div class="">{{userStore.user.phone}}</div>
    <div class="">{{userStore.user.address}}</div>
    <div class="">{{userStore.user.username}}</div>
  </main>
  <Footer />
</template>
