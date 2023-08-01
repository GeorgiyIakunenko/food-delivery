<script setup>
import { useUserStore } from "@/store/user";
import Modal from "@/components/Modal.vue";
import { ref } from "vue";

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
  <main>
    <div class="container">
      Home

      <button @click="openModal">Open modal</button>
      <Modal :type="modalType" @modalClose="closeModal" v-bind:open="modalOpen"
        >Error</Modal
      >
    </div>
  </main>
</template>

<style scoped></style>
