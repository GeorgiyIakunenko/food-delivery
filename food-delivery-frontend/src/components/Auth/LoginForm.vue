<script setup>
import { RouterLink, RouterView } from 'vue-router'
import {defineProps, computed, reactive} from 'vue';

const props = defineProps({
  toggleSingIn: {
    type: Function,
    required: true
  },
});

const activeInputs = reactive({});

const toggleFocus = (inputKey) => {
  activeInputs[inputKey] = !activeInputs[inputKey];
};

const isActive = (inputKey) => {
  return activeInputs[inputKey] || false;
}

</script>

<template>
  <form action="#" autocomplete="off" class="sign-in-form">
    <div class="heading">
      <h2>Welcome Back</h2>
      <h6>Not registered yet?</h6>
      <a href="#" @click="toggleSingIn"  class="toggle"   >Sign up</a>
    </div>

    <div class="actual-form">
      <div  class="input-wrap">
        <input
            @focus="toggleFocus('loginName')"
            @blur="toggleFocus('loginName')"
            :class="{'active': isActive('loginName')}"
            type="text"
            minlength="4"
            class="input-field"
            autocomplete="off"
            required
        />
        <label>Name</label>
      </div>

      <div class="input-wrap">
        <input
            @focus="toggleFocus('loginPassword')"
            @blur="toggleFocus('loginPassword')"
            :class="{'active': isActive('loginPassword')}"
            type="password"
            minlength="4"
            class="input-field"
            autocomplete="off"
            required
        />
        <label>Password</label>
      </div>

      <input type="submit" value="Sign In" class="sign-btn" />

      <p class="text">
        Forgotten your password or you login details?
        <a class="reset-password-btn" href="#">Reset Password</a>
      </p>
    </div>
  </form>
</template>

<style scoped>
form {
  max-width: 260px;
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

form.sign-up-form {
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
.sign-btn {
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

.sign-btn:hover {
  background-color: #8371fd;
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

.reset-password-btn {
  display: inline-block;
  color: #151111!important;
  text-decoration: none;
  font-size: 0.75rem;
  font-weight: 500;
  transition: 0.3s;
}

.reset-password-btn:hover {
  color: #8371fd!important;
}
</style>