<script setup>
import Footer from "@/components/Footer.vue";
import Header from "@/components/Header.vue";
import { useUserStore } from "@/stores/user";
import { onMounted, ref } from "vue";
import {getUserData, logout, updateProfile} from "@/api/api";
import Button from "@/components/UI/Button.vue";

const userStore = useUserStore();

onMounted(() => {
  getUserData();
});

const updateUserProfile = () => {
  const updateUserRequestBody = {
    first_name: userStore.updatedUser.first_name !== userStore.user.first_name ? userStore.updatedUser.first_name : "",
    last_name: userStore.updatedUser.last_name !== userStore.user.last_name ? userStore.updatedUser.last_name : "",
    username: "",
    age: userStore.updatedUser.age !== userStore.user.age ? userStore.updatedUser.age : 0,
    phone: userStore.updatedUser.phone !== userStore.user.phone ? userStore.updatedUser.phone : "",
    address: userStore.updatedUser.address !== userStore.user.address ? userStore.updatedUser.address : "",
  };

  console.log(updateUserRequestBody)

  updateProfile(updateUserRequestBody);
};


const isEditMode = ref(false);

const toggleEditMode = () => {
  isEditMode.value = !isEditMode.value;
  if (isEditMode.value) {
    userStore.userUpdate(userStore.user)
  }
};
</script>

<template>
  <Header />
  <main class="profile-container">
    <div class="container">
      <div class="user-data">
        <h1 class="profile-heading">Profile</h1>
        <div v-if="isEditMode">
          <form @submit.prevent="updateUserProfile" class="profile-form">
            <div class="profile-form-group">
              <label class="profile-form-label" for="name">First Name</label>
              <input v-model="userStore.updatedUser.first_name" class="profile-form-input" type="text" id="name" required />
            </div>
            <div class="profile-form-group">
              <label class="profile-form-label" for="name">Last Name</label>
              <input v-model="userStore.updatedUser.last_name" class="profile-form-input" type="text" id="name" required />
            </div>
            <div class="profile-form-group">
              <label class="profile-form-label" for="age">Age</label>
              <input v-model="userStore.updatedUser.age" class="profile-form-input" type="number" id="age" required />
            </div>
            <div class="profile-form-group">
              <label class="profile-form-label" for="phone">Phone</label>
              <input v-model="userStore.updatedUser.phone" class="profile-form-input" type="tel" id="phone" required />
            </div>
            <div class="profile-form-group">
              <label class="profile-form-label" for="address">Address</label>
              <input v-model="userStore.updatedUser.address" class="profile-form-input" type="text" id="address" required />
            </div>
<!--            <div class="profile-form-group">
              <label class="profile-form-label" for="username">Username</label>
              <input v-model="userStore.updatedUser.username" class="profile-form-input" type="text" id="username" required />
            </div>-->
            <button  @click="updateUserProfile" type="submit" class="profile-form-button">Save</button>
          </form>
        </div>
        <div v-else>
          <div class="profile-info">
            <p><strong>First Name</strong> {{ userStore.user.first_name }}</p>
            <p><strong>Last Name</strong> {{ userStore.user.last_name }}</p>
            <p><strong>Email:</strong> {{ userStore.user.email }}</p>
            <p><strong>Age:</strong> {{ userStore.user.age }}</p>
            <p><strong>Phone:</strong> {{ userStore.user.phone }}</p>
            <p><strong>Address:</strong> {{ userStore.user.address }}</p>
            <p><strong>Username:</strong> {{ userStore.user.username }}</p>
            <router-link to="/products" @click="logout" ><Button class="logout-btn"  intent="secondary">Logout</Button> </router-link>

          </div>
        </div>
        <button @click="toggleEditMode" class="edit-button">{{ isEditMode ? 'Cancel' : 'Edit' }}</button>

      </div>
      <div class="user-orders">

      </div>
    </div>

  </main>
  <Footer />
</template>

<style>

.logout-btn {
  margin-top: 20px;
}

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

.profile-info p {
  margin-bottom: 1rem;
}

.profile-form {
  display: grid;
  grid-gap: 1rem;
  max-width: 50%;
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


@media (max-width: 768px) {
  .profile-form {
    max-width: 100%;
  }
}




</style>
