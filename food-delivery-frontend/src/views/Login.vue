<script setup>
  import { RouterLink, RouterView } from 'vue-router'
  import {useUserStore} from "../stores/user";
  import Header from "@/components/Header.vue";
  import Footer from "@/components/Footer.vue";

  import {ref, reactive} from "vue";
  const singIn = ref(true)

  const toggleSingIn = () => {
    singIn.value = !singIn.value
  }
  const userStore = useUserStore()

  const activeInputs = reactive({});

  const toggleFocus = (inputKey) => {
    activeInputs[inputKey] = !activeInputs[inputKey];
  };

  const isActive = (inputKey) => {
    return activeInputs[inputKey] || false;
  }
</script>

<template>
  <Header></Header>
  <main :class="{'sign-up-mode': !singIn}">
    <div class="box">
      <div class="inner-box">
        <div class="forms-wrap">
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
          <form action="#" autocomplete="off" class="sign-up-form">


            <div class="heading">
              <h2>Get Started</h2>
              <h6>Already have an account?</h6>
              <a href="#" @click="toggleSingIn" class="toggle">Sign in</a>
            </div>

            <div class="actual-form">
              <div class="input-wrap">
                <input
                    @focus="toggleFocus('registerName')"
                    @blur="toggleFocus('registerName')"
                    :class="{'active': isActive('registerName')}"
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

              <input type="submit" value="Sign Up" class="sign-btn" />

              <p class="text">
                By signing up, I agree to the
                <a href="#">Terms of Services</a> and
                <a href="#">Privacy Policy</a>
              </p>
            </div>
          </form>
        </div>

        <div class="carousel">
          <div class="images-wrapper">
            <img src="@/assets/images/icons/login.png" class="image img-1" :class="{show : singIn}" alt="" />
            <img src="@/assets/images/icons/register.png" class="image img-2" :class="{show : !singIn}" alt="" />
          </div>
        </div>
      </div>
    </div>
  </main>
  <Footer></Footer>
</template>

<style scoped>

main {
  background-color: #aaaaaa;
  display: flex;
  align-items: center;
}

.box {
  position: relative;
  width: 95%;
  margin: 0 auto;
  max-width: 1020px;
  height: 640px;
  background-color: #fff;
  border-radius: 3.3rem;
  box-shadow: 0 60px 40px -30px rgba(0, 0, 0, 0.27);
}

.toggle {
  padding: 0 15px;
}

.inner-box {
  position: absolute;
  width: calc(100% - 2rem);
  height: calc(100% - 2rem);
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.forms-wrap {
  position: absolute;
  height: 100%;
  width: 45%;
  top: 0;
  left: 0;
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: 1fr;
  transition: 0.8s ease-in-out;
}

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

main.sign-up-mode form.sign-in-form {
  opacity: 0;
  pointer-events: none;
}

main.sign-up-mode form.sign-up-form {
  opacity: 1;
  pointer-events: all;
}

main.sign-up-mode .forms-wrap {
  left: 55%;
}

main.sign-up-mode .carousel {
  left: 0;
}

.carousel {
  position: absolute;
  height: 100%;
  width: 55%;
  left: 45%;
  top: 0;
  background-color: #ffe0d2;
  border-radius: 2rem;
  display: grid;
  grid-template-rows: auto 1fr;
  padding-bottom: 2rem;
  overflow: hidden;
  transition: 0.8s ease-in-out;
}

.images-wrapper {
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: 1fr;
}

.image {
  width: 100%;
  grid-column: 1/2;
  grid-row: 1/2;
  opacity: 0;
  transition: opacity 0.3s, transform 0.5s;
}

.img-1 {
  transform: translate(0, -50px);
}

.img-2 {
  transform: scale(0.4, 0.5);
}

.image.show {
  opacity: 1;
  transform: none;
}

@media (max-width: 850px) {
  .box {
    height: auto;
    max-width: 550px;
    overflow: hidden;
  }

  .inner-box {
    position: static;
    transform: none;
    width: revert;
    height: revert;
    padding: 2rem;
  }

  .forms-wrap {
    position: revert;
    width: 100%;
    height: auto;
  }

  form {
    max-width: revert;
    padding: 1.5rem 2.5rem 2rem;
    transition: transform 0.8s ease-in-out, opacity 0.45s linear;
  }

  .heading {
    margin: 2rem 0;
  }

  form.sign-up-form {
    transform: translateX(100%);
  }

  main.sign-up-mode form.sign-in-form {
    transform: translateX(-100%);
  }

  main.sign-up-mode form.sign-up-form {
    transform: translateX(0%);
  }

  .carousel {
    display: none;
  }

  .images-wrapper {
    display: none;
  }
}

@media (max-width: 530px) {
  main {
    padding: 1rem;
  }

  .box {
    border-radius: 2rem;
  }

  .inner-box {
    padding: 1rem;
  }

  .carousel {
    padding: 1.5rem 1rem;
    border-radius: 1.6rem;
  }

  form {
    padding: 1rem 2rem 1.5rem;
  }
}
</style>