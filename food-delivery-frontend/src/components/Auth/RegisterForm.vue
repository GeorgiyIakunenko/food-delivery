<script setup>
  import {reactive} from "vue";

  const activeInputs = reactive({});

  const toggleFocus = (inputKey) => {
    activeInputs[inputKey] = !activeInputs[inputKey];
  };
  const isActive = (inputKey) => {
    return activeInputs[inputKey] || false;
  }
</script>
<template>
  <form action="#" autocomplete="off" class="sign-up-form">
    <div class="heading">
      <h2>Get Started</h2>
      <h6>Already have an account?</h6>
      <a href="#" @click="$emit('changeMode')" class="toggle">Sign in</a>
    </div>
    <div class="actual-form">
      <div class="actual-form_content">
        <div class="input-wrap">
          <input
              @focus="toggleFocus('registerFirstName')"
              @blur="toggleFocus('registerFirstName')"
              :class="{'active': isActive('registerFirstName')}"
              type="text"
              minlength="4"
              class="input-field"
              autocomplete="off"
              required
          />
          <label>First name</label>
        </div>
        <div class="input-wrap">
          <input
              @focus="toggleFocus('registerLastName')"
              @blur="toggleFocus('registerLastName')"
              :class="{'active': isActive('registerLastName')}"
              type="text"
              minlength="4"
              class="input-field"
              autocomplete="off"
              required
          />
          <label>Last name</label>
        </div>
        <div class="input-wrap">
          <input
              @focus="toggleFocus('registerEmail')"
              @blur="toggleFocus('registerEmail')"
              :class="{'active': isActive('registerEmail')}"
              type="email"
              class="input-field"
              autocomplete="off"
              required
          />
          <label>Email</label>
        </div>
        <div class="input-wrap">
          <input
              @focus="toggleFocus('registerUsername')"
              @blur="toggleFocus('registerUsername')"
              :class="{'active': isActive('registerUsername')}"
              type="password"
              minlength="4"
              class="input-field"
              autocomplete="off"
              required
          />
          <label>Username</label>
        </div>
        <div class="input-wrap">
          <input
              @focus="toggleFocus('registerAge')"
              @blur="toggleFocus('registerAge')"
              :class="{'active': isActive('registerAge')}"
              type="password"
              minlength="4"
              class="input-field"
              autocomplete="off"
              required
          />
          <label>Age</label>
        </div>
        <div class="input-wrap">
          <input
              @focus="toggleFocus('registerAddress')"
              @blur="toggleFocus('registerAddress')"
              :class="{'active': isActive('registerAddress')}"
              type="password"
              minlength="4"
              class="input-field"
              autocomplete="off"
              required
          />
          <label>Address</label>
        </div>
        <div class="input-wrap">
          <input
              @focus="toggleFocus('registerPassword')"
              @blur="toggleFocus('registerPassword')"
              :class="{'active': isActive('registerPassword')}"
              type="password"
              minlength="4"
              class="input-field"
              autocomplete="off"
              required
          />
          <label>Password</label>
        </div>
      </div>
      <input type="submit" value="Sign Up" class="sign-btn" />
      <p class="text">
        By signing up, I agree to the
        <a href="#">Terms of Services</a> and
        <a href="#">Privacy Policy</a>
      </p>
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


main.sign-up-mode form.sign-up-form {
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

  main.sign-up-mode form {
    transform: translateX(0%);
  }
}

@media (max-width: 530px) {
  main {
    padding: 1rem;
  }

}
</style>