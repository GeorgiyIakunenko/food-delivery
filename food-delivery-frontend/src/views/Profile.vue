<script setup>
import Footer from "@/components/Footer.vue";
import Header from "@/components/Header.vue";
import { useUserStore } from "@/stores/user";
import { onMounted, ref } from "vue";
import { getUserData } from "@/api/api";

const userStore = useUserStore();

onMounted(() => {
  getUserData();
});

const isEditMode = ref(false);

const toggleEditMode = () => {
  isEditMode.value = !isEditMode.value;
};

const updateUserProfile = () => {
  // Implement the logic to update the user profile data
  // For this example, let's just show an alert to indicate success
  alert("Profile updated successfully!");
};
</script>

<template>
  <Header />
  <main class="profile-container">
    <div class="user-data">
      <h1 class="profile-heading">Profile</h1>
      <div v-if="isEditMode">
        <form @submit.prevent="updateUserProfile" class="profile-form">
          <div class="profile-form-group">
            <label class="profile-form-label" for="name">Name</label>
            <input v-model="userStore.user.name" class="profile-form-input" type="text" id="name" required />
          </div>
          <div class="profile-form-group">
            <label class="profile-form-label" for="email">Email</label>
            <input v-model="userStore.user.email" class="profile-form-input" type="email" id="email" required />
          </div>
          <div class="profile-form-group">
            <label class="profile-form-label" for="age">Age</label>
            <input v-model="userStore.user.age" class="profile-form-input" type="number" id="age" required />
          </div>
          <div class="profile-form-group">
            <label class="profile-form-label" for="phone">Phone</label>
            <input v-model="userStore.user.phone" class="profile-form-input" type="tel" id="phone" required />
          </div>
          <div class="profile-form-group">
            <label class="profile-form-label" for="address">Address</label>
            <input v-model="userStore.user.address" class="profile-form-input" type="text" id="address" required />
          </div>
          <div class="profile-form-group">
            <label class="profile-form-label" for="username">Username</label>
            <input v-model="userStore.user.username" class="profile-form-input" type="text" id="username" required />
          </div>
          <button type="submit" class="profile-form-button">Save</button>
        </form>
      </div>
      <div v-else>
        <div class="profile-info">
          <p><strong>Name:</strong> {{ userStore.user.name }}</p>
          <p><strong>Email:</strong> {{ userStore.user.email }}</p>
          <p><strong>Age:</strong> {{ userStore.user.age }}</p>
          <p><strong>Phone:</strong> {{ userStore.user.phone }}</p>
          <p><strong>Address:</strong> {{ userStore.user.address }}</p>
          <p><strong>Username:</strong> {{ userStore.user.username }}</p>
        </div>
      </div>
      <button @click="toggleEditMode" class="edit-button">{{ isEditMode ? 'Cancel' : 'Edit' }}</button>
    </div>
    <div class="user-orders">

    </div>
  </main>
  <Footer />
</template>

<style>

.profile-container {
  width: 90%;
  margin: 90px auto 15px;
  padding: 2rem;

  border-radius: 8px;
  box-shadow: 1px 1px 8px 2px rgba(0, 0, 0, 0.12);
}

.profile-heading {
  text-align: center;
  font-size: 1.5rem;
  margin-bottom: 2rem;
  color: #333;
}

.profile-info {
  font-size: 1rem;
  color: #333;
  line-height: 1.5;
}

.profile-form {
  display: grid;
  grid-gap: 1rem;
}

.profile-form-label {
  font-size: 1rem;
  font-weight: bold;
  color: #333;
}

.profile-form-input {
  width: 100%;
  padding: 0.5rem;
  font-size: 1rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.profile-form-button {
  width: 100%;
  padding: 0.75rem 1rem;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: bold;
  transition: background-color 0.3s ease;
}

.profile-form-button:hover {
  background-color: #0056b3;
}

.edit-button {
  padding: 0.75rem 1rem;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: bold;
  transition: background-color 0.3s ease;
  margin-top: 1rem;
}

.edit-button:hover {
  background-color: #0056b3;
}
</style>
