<script setup>
import { useUserStore } from "@/store/user";
import Input from "@/components/Input.vue";
import Button from "@/components/Button.vue";
import { computed, reactive, ref } from "vue";
import router from "@/router/router";
import { getProfile, updateProfile } from "@/api/user";
import Modal from "@/components/Modal.vue";

const userStore = useUserStore();

const userChangeForm = reactive({
  FirstName: userStore.user.FirstName,
  LastName: userStore.user.LastName,
  Username: userStore.user.Username,
  Email: userStore.user.Email,
  Age: userStore.user.Age,
  Address: userStore.user.Address,
});

const isButtonDisabled = computed(() => {
  return (
    userChangeForm.FirstName === userStore.user.FirstName &&
    userChangeForm.LastName === userStore.user.LastName &&
    userChangeForm.Username === userStore.user.Username &&
    userChangeForm.Age === userStore.user.Age &&
    userChangeForm.Address === userStore.user.Address
  );
});

let modalTitle = "success";
let modalMessage = "We send the code to your email address ";
let modalType = "success";
let modalOpen = ref(false);

const openWarningModal = () => {
  modalTitle = "Warning";
  modalMessage = "Are you sure you want to change your profile information?";
  modalType = "warning";
  modalOpen.value = true;
};

const submitForm = async (isSubmit) => {
  if (!isSubmit || modalType === "success" || modalType === "error") {
    modalOpen.value = false;
    return;
  }

  modalOpen.value = false;
  const changeUserForm = {
    first_name:
      userStore.user.FirstName !== userChangeForm.FirstName
        ? userChangeForm.FirstName
        : "",
    last_name:
      userStore.user.LastName !== userChangeForm.LastName
        ? userChangeForm.LastName
        : "",
    username:
      userStore.user.Username !== userChangeForm.Username
        ? userChangeForm.Username
        : "",
    age: userStore.user.Age !== userChangeForm.Age ? userChangeForm.Age : 0,
    phone: "",
    address:
      userStore.user.Address !== userChangeForm.Address
        ? userChangeForm.Address
        : "",
  };

  const res = await updateProfile(changeUserForm);
  if (res.success) {
    modalTitle = "success";
    modalMessage = "Your profile data has been changed";
    modalType = "success";
    modalOpen.value = true;
  } else {
    modalTitle = "Error";
    modalMessage = "Something went wrong";
    modalType = "error";
    modalOpen.value = true;
  }
};

const closeChangeMood = async () => {
  isChangeMood.value = false;
  await getProfile();
};

const isChangeMood = ref(false);
</script>
<template>
  <main>
    <div class="container">
      <div class="font-sans">
        <div v-if="!isChangeMood" class="mb-10 text-center text-2xl font-bold">
          Profile Information
        </div>
        <div v-else class="mb-10 text-center text-2xl font-bold text-amber-700">
          Change profile information
        </div>
        <div class="mx-auto flex w-4/5 flex-col gap-7 md:w-3/5">
          <Input
            label="First Name"
            name="First Name"
            v-model="userChangeForm.FirstName"
            :disabled="!isChangeMood"
          ></Input>
          <Input
            label="Last Name"
            name="Last Name"
            v-model="userChangeForm.LastName"
            :disabled="!isChangeMood"
          ></Input>
          <Input
            v-if="!isChangeMood"
            label="Username"
            name="Username"
            v-model="userChangeForm.Username"
            :disabled="!isChangeMood"
          ></Input>
          <Input
            v-if="!isChangeMood"
            label="Email"
            name="Email"
            v-model="userChangeForm.Email"
            :disabled="!isChangeMood"
          ></Input>
          <Input
            label="Age"
            name="Age"
            v-model="userChangeForm.Age"
            :disabled="!isChangeMood"
          ></Input>
          <Input
            label="Address"
            name="Address"
            v-model="userChangeForm.Address"
            :disabled="!isChangeMood"
          ></Input>
          <div v-if="!isChangeMood" class="flex flex-col gap-4">
            <Button
              class="w-full"
              type="primary"
              @click="isChangeMood = !isChangeMood"
              >Edit</Button
            >
            <Button
              class="w-full"
              type="secondary"
              @click="$router.push('/profile')"
              >Back</Button
            >
          </div>
          <div v-else class="flex gap-3">
            <Button class="w-full" type="secondary" @click="closeChangeMood"
              >Cancel</Button
            >
            <Button
              class="w-full"
              :disabled="isButtonDisabled"
              type="primary"
              @click="openWarningModal"
              >Save</Button
            >
          </div>
        </div>
      </div>
    </div>
    <Modal
      :type="modalType"
      :title="modalTitle"
      @modalClose="submitForm"
      v-bind:open="modalOpen"
      >{{ modalMessage }}</Modal
    >
  </main>
</template>

<style scoped></style>
