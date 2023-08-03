<script setup>
import { useUserStore } from "@/store/user";
import { onMounted, ref } from "vue";
import Filter from "@/components/Filter/Filter.vue";
import { getProducts } from "@/api/api";

onMounted(() => {
  getProducts();
});

const userStore = useUserStore();

const modalOpen = ref(false);
const modalType = ref("error");
const openModal = () => {
  modalOpen.value = true;
};
const closeModal = (value) => {
  modalOpen.value = false;
  alert(value);
};
</script>

<template>
  <main class="pt-24">
    <div class="container">
      <!--      Home

      <button @click="openModal">Open modal</button>
      <Modal
        :type="modalType"
        title="Failed to register"
        @modalClose="closeModal"
        v-bind:open="modalOpen"
        >Password is not correct</Modal
      >-->
      <Filter></Filter>
      <div class="">Content</div>
    </div>
  </main>
</template>

<style scoped></style>
