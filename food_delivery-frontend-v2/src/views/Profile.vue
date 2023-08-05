<script setup>
import Button from "@/components/Button.vue";
import { logout } from "@/api/api";
import { useRouter } from "vue-router";
import Modal from "@/components/Modal.vue";
import { ref } from "vue";

const router = useRouter();

let modalTitle = "Are you sure you want to logout?";
let modalMessage = "You will be redirected to the login page.";
let modalType = "warning";
let modalOpen = ref(false);

const closeModal = (value) => {
  console.log("close modal", value);
  modalOpen.value = false;

  if (value === true) {
    logoutUser();
  }
};

const logoutUser = async () => {
  const result = await logout();
  if (result.success === true) {
    console.log(result);
    await router.push("/login");
  } else {
    console.log(result);
    alert("Failed to logout");
  }
};
</script>

<template>
  <main>
    <div class="container">
      <div class="">Profile</div>
      <Button @click="modalOpen = true">Logout</Button>
    </div>
    <Modal
      :type="modalType"
      :title="modalTitle"
      @modalClose="closeModal"
      v-bind:open="modalOpen"
      >{{ modalMessage }}</Modal
    >
  </main>
</template>

<style scoped></style>
