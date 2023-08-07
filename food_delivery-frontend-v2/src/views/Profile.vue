<script setup>
import Button from "@/components/Button.vue";
import { getProfile, logout } from "@/api/user";
import { useRouter } from "vue-router";
import Modal from "@/components/Modal.vue";
import { onMounted, ref } from "vue";
import { UserCircleIcon } from "@heroicons/vue/24/outline";
import { useUserStore } from "@/store/user";

const router = useRouter();
const userStore = useUserStore();

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

onMounted(() => {
  const res = getProfile();
  console.log(res);
});
</script>

<template>
  <main>
    <div class="container font-sans">
      <div class="mb-4 text-center text-xl font-medium text-neutral-800">
        Profile
      </div>
      <div
        class="flex flex-col items-center justify-center border-b-2 border-t-2 py-4"
      >
        <UserCircleIcon class="mb-3 h-20 w-20 text-neutral-500" />
        <div class="">
          {{ userStore.user.FirstName }} {{ userStore.user.LastName }}
        </div>
      </div>
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
