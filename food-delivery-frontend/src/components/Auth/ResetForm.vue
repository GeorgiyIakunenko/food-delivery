<script setup>

import {reactive, ref} from "vue";
import {useUserStore} from "@/stores/user";
import {ResetPassword} from "@/api/api";

const activeInputs = reactive({});

const userStore = useUserStore()

const toggleFocus = (inputKey) => {
  activeInputs[inputKey] = !activeInputs[inputKey];
}

const isActive = (inputKey) => {
  return activeInputs[inputKey] || false;
}

const emit = defineEmits(['changeMode'])

const dialogMessage = ref('you successfully reset your password')
const dialogType = ref('success')


const showDialog = () => {
  emit('showModal', {
    type: dialogType.value,
    message: dialogMessage.value
  })
}

const SubmitForm = () => {

  ResetPassword(userStore.resetForm.email, userStore.resetForm.reset_code, userStore.resetForm.new_password).then((response) => {
    if (response) {
      dialogType.value = 'success'
      dialogMessage.value = 'you successfully reset your password'
      emit('changeMode')
      console.log(response)
      showDialog()
    } else {
      dialogType.value = 'error'
      dialogMessage.value = 'You enter wrong code'
      showDialog()
    }
  })
}

</script>

<template>
  <form action="#" autocomplete="off" class="sign-up-form form-2">
    <div class="heading">
      <button @click="$emit('changeMode')" class="toggle auth-btn">Go back</button>
      <h2>Create a New Password</h2>
      <h6>Enter the verification code and set your new password</h6>
    </div>
    <div class="actual-form">
      <div class="actual-form_content">
        <div class="input-wrap">
          <input
              @focus="toggleFocus('resetCode')"
              @blur="toggleFocus('resetCode')"
              :class="{'active': isActive('resetCode') || userStore.resetForm.reset_code !== ''}"
              v-model="userStore.resetForm.reset_code"
              type="text"
              class="input-field"
              autocomplete="off"
              required
          />
          <label>Code from email</label>
        </div>
         <div class="input-wrap">
          <input
              @focus="toggleFocus('resetPassword')"
              @blur="toggleFocus('resetPassword')"
              :class="{'active': isActive('resetPassword') || userStore.resetForm.new_password !== ''}"
              v-model="userStore.resetForm.new_password"
              type="password"
              minlength="6"
              class="input-field"
              autocomplete="off"
              required
          />
          <label>New password</label>
        </div>
      </div>
      <input @click.stop.prevent="SubmitForm" type="submit" value="Reset Password" class="auth-btn" />
    </div>
  </form>
</template>

<style scoped>
.toggle {
  padding: 0 15px;
}

form {
  max-width: 320px;
  width: 100%;
  margin: 0 auto;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  grid-column: 1 / 2;
  grid-row: 1 / 2;
  transition: opacity 0.02s 0.4s;
}

form.form-2 {
  opacity: 0;
  pointer-events: none;
}

.heading h2 {
  font-size: 2.1rem;
  font-weight: 600;
  color: #151111;
}

.heading h6 {
  color: #bababa;
  font-weight: 400;
  font-size: 0.75rem;
  display: inline;
}

.toggle {
  color: #151111;
  text-decoration: none;
  font-size: 0.75rem;
  font-weight: 500;
  transition: 0.3s;
}

.toggle:hover {
  color: #8371fd;
}

.input-wrap {
  position: relative;
  height: 37px;
  margin-bottom: 2rem;
}

.input-field {
  position: absolute;
  width: 100%;
  height: 100%;
  background: none;
  border: none;
  outline: none;
  border-bottom: 1px solid #bbb;
  padding: 0;
  font-size: 0.95rem;
  color: #151111;
  transition: 0.4s;
}

label {
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  font-size: 0.95rem;
  color: #bbb;
  pointer-events: none;
  transition: 0.4s;
}

.input-field.active {
  border-bottom-color: #151111;
}

.input-field.active + label {
  font-size: 0.75rem;
  top: -2px;
}
.auth-btn {
  display: inline-block;
  width: 100%;
  height: 43px;
  background-color: #151111;
  color: #fff;
  border: none;
  cursor: pointer;
  border-radius: 0.8rem;
  font-size: 0.8rem;
  margin-bottom: 2rem;
  transition: 0.3s;
}

.auth-btn:hover {
  background-color: #8371fd;
  color: white;
}

.text {
  color: #bbb;
  font-size: 0.7rem;
}

.text a {
  color: #bbb;
  transition: 0.3s;
}

.text a:hover {
  color: #8371fd;
}

.actual-form_content {
  padding: 15px;
  max-height: 400px;
  overflow: auto;
}

.actual-form_content::-webkit-scrollbar {
  width: 8px;
}

.actual-form_content::-webkit-scrollbar-track {
  background-color: #f1f1f1;
}

.actual-form_content::-webkit-scrollbar-thumb {
  background-color: #888;
}

.actual-form_content::-webkit-scrollbar-thumb:hover {
  background-color: #555;
}


main.form-2-mode form.form-2 {
  opacity: 1;
  pointer-events: all;
}

@media (max-width: 850px) {

  form {
    max-width: revert;
    transition: transform 0.8s ease-in-out, opacity 0.45s linear;
    transform: translateX(100%);
  }

  .heading {
    margin: 2rem 0;
  }

  main.form-2-mode form {
    transform: translateX(0%);
  }
}

@media (max-width: 530px) {
  main {
    padding: 1rem;
  }

}
</style>