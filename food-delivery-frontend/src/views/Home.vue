<script setup>
import Footer from "@/components/Footer.vue";
import Header from "@/components/Header.vue";
import {useUserStore} from "@/stores/user";
import Dialog from "@/components/UI/Dialog.vue";
import {ref} from "vue";

const userStore = useUserStore()

const dialogVisible = ref(false)

const toggleDialog = () => {
  dialogVisible.value = !dialogVisible.value
}

</script>

<template>
  <Header />
  <main>
    HomePage
    <button @click="toggleDialog">Show Dialog</button>
    <Dialog @update:show="toggleDialog" type="error" v-model:show="dialogVisible"> gsdgadsfds afdasfasdf asdfasdfas dfasfsda fsdfasd fasdfasfasfssdf </Dialog>
  </main>

  <Footer />
</template>
