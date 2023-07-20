<script setup>
import { RouterLink, RouterView } from 'vue-router'
import {defineProps, computed, reactive} from 'vue';


const ResetRequestForm = reactive({
  email: '',
})

const activeInputs = reactive({});

const toggleFocus = (inputKey) => {
  activeInputs[inputKey] = !activeInputs[inputKey];
};
const isActive = (inputKey) => {
  return activeInputs[inputKey] || false;
}

</script>

<template>
  <form @submit="" action="#" autocomplete="off" class="reset-password-request form-1">
    <div class="heading">
      <button  @click="$emit('change-reset-mode')" class="toggle auth-btn">Go back</button>
      <h2>Forgot your password? <br> No problem.</h2>
      <h6>We'll send a verification code to your email address, allowing you to reset your password.</h6>
    </div>
    <div class="actual-form">
      <div  class="input-wrap">
        <input
            @focus="toggleFocus('resetEmail')"
            @blur="toggleFocus('resetEmail')"
            :class="{'active': isActive('resetEmail')}"
            v-model="ResetRequestForm.email"
            type="email"
            minlength="6"
            class="input-field"
            autocomplete="off"
            required
        />
        <label>Email</label>
      </div>
      <input type="submit" @click="$emit('changeMode')" value="Send Code" class="auth-btn" />
    </div>
  </form>
</template>
<style scoped>

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

.heading h2 {
  font-size: 1.5rem;
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
  padding: 0 15px;
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

main.form-2-mode form.form-1 {
  opacity: 0;
  pointer-events: none;
}

@media (max-width: 850px) {

  form {
    max-width: revert;
    padding: 1.5rem 2.5rem 2rem;
    transition: transform 0.8s ease-in-out, opacity 0.45s linear;
  }

  .heading {
    margin: 2rem 0;
  }
}

@media (max-width: 530px) {
  main {
    padding: 1rem;
  }

}

</style>