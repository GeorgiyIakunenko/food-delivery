<script setup>
import Button from "@/components/Button.vue";
import { getProfile, logout } from "@/api/user";
import { useRouter } from "vue-router";
import Modal from "@/components/Modal.vue";
import { onMounted, ref } from "vue";
import {
  UserCircleIcon,
  ChevronRightIcon,
  ArrowRightOnRectangleIcon,
} from "@heroicons/vue/24/outline";
import { useUserStore } from "@/store/user";

const router = useRouter();
const userStore = useUserStore();

let modalTitle = "Are you sure you want to logout?";
let modalMessage = "You will be redirected to the login page.";
let modalType = "warning";
let modalOpen = ref(false);

let pressedButton = "";

const closeModal = (value) => {
  console.log("close modal", value);
  modalOpen.value = false;

  if (value === true && pressedButton === "logout") {
    logoutUser();
  }
};

const logoutButton = () => {
  modalTitle = "Are you sure you want to logout?";
  modalMessage = "You will be redirected to the login page.";
  modalType = "warning";
  modalOpen.value = true;
  pressedButton = "logout";
};

const logoutUser = async () => {
  const result = await logout();
  if (result.success === true) {
    console.log(result);
    await router.push("/login");
  } else {
    console.log("failed to logout");
  }
};

onMounted(() => {
  const res = getProfile();
  console.log(res);
});
</script>

<template>
  <main class="bg-card-bg">
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
      <div class="mt-10 flex flex-col gap-5">
        <button
          @click="$router.push('/profile/orders')"
          class="mx-auto flex w-4/5 items-center rounded-xl bg-neutral-0 px-5 py-3 transition-all hover:bg-neutral-30"
        >
          <ArrowRightOnRectangleIcon
            class="mr-3 h-5 w-5 text-neutral-200"
          ></ArrowRightOnRectangleIcon>
          <span>Orders</span>

          <ChevronRightIcon
            class="ml-auto h-5 w-5 text-neutral-200"
          ></ChevronRightIcon>
        </button>
        <button
          @click="$router.push('/profile-info')"
          class="mx-auto flex w-4/5 items-center rounded-xl bg-neutral-0 px-5 py-3 transition-all hover:bg-neutral-30"
        >
          <ArrowRightOnRectangleIcon
            class="mr-3 h-5 w-5 text-neutral-200"
          ></ArrowRightOnRectangleIcon>
          <span>Profile Info</span>
          <ChevronRightIcon
            class="ml-auto h-5 w-5 text-neutral-200"
          ></ChevronRightIcon>
        </button>
        <button
          @click="$router.push('/profile-password')"
          class="mx-auto flex w-4/5 items-center rounded-xl bg-neutral-0 px-5 py-3 transition-all hover:bg-neutral-30"
        >
          <ArrowRightOnRectangleIcon
            class="mr-3 h-5 w-5 text-neutral-200"
          ></ArrowRightOnRectangleIcon>
          <span>Change Password</span>

          <ChevronRightIcon
            class="ml-auto h-5 w-5 text-neutral-200"
          ></ChevronRightIcon>
        </button>
        <button
          @click="logoutButton"
          class="mx-auto flex w-4/5 items-center rounded-xl bg-neutral-0 px-5 py-3 transition-all hover:bg-neutral-30"
        >
          <ArrowRightOnRectangleIcon
            class="mr-3 h-5 w-5 text-neutral-200"
          ></ArrowRightOnRectangleIcon>
          <span>Log Out</span>

          <ChevronRightIcon
            class="ml-auto h-5 w-5 text-neutral-200"
          ></ChevronRightIcon>
        </button>
      </div>
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
