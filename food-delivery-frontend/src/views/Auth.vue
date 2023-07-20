<script setup>
  import { RouterLink, RouterView } from 'vue-router'
  import {useUserStore} from "../stores/user";
  import Header from "@/components/Header.vue";
  import Footer from "@/components/Footer.vue";
  import LoginForm from "@/components/Auth/LoginForm.vue";
  import RegisterForm from "@/components/Auth/RegisterForm.vue";
  import {ref, reactive} from "vue";
  import Dialog from "@/components/UI/Dialog.vue";
  import ResetForm from "@/components/Auth/ResetForm.vue";
  import ResetRequestForm from "@/components/Auth/ResetRequestForm.vue";

  
  const changeMode = ref(true)
  const showDialog = ref(false)

  const resetMode = ref(false)

  const toggleDialog = () => {
    showDialog.value = !showDialog.value
  }
  const toggleChangeMode = () => {
    changeMode.value = !changeMode.value
  }
  const toggleReset = () => {
    resetMode.value = !resetMode.value
  }
</script>

<template>
  <Header></Header>
  <main :class="{'form-2-mode': !changeMode, 'reset-mode' : resetMode }">
    <div class="box">
      <div class="inner-box">
        <div v-if="!resetMode" class="forms-wrap">
          <LoginForm @reset-password="toggleReset"  @changeMode="toggleChangeMode"></LoginForm>
          <RegisterForm @changeMode="toggleChangeMode" ></RegisterForm>
        </div>
        <div v-else class="forms-wrap">
          <ResetRequestForm  @change-reset-mode="toggleReset"  @changeMode="toggleChangeMode"></ResetRequestForm>
          <ResetForm  @changeMode="toggleChangeMode"></ResetForm>
        </div>
        <div class="carousel">
          <div v-if="!resetMode" class="images-wrapper">
            <img src="@/assets/images/icons/login.png" class="image img-login" :class="{show : changeMode}" alt="" />
            <img src="@/assets/images/icons/register.png" class="image img-register" :class="{show : !changeMode}" alt="" />
          </div>
          <div v-else class="images-wrapper">
            <img src="@/assets/images/auth/reset-request.png" class="image img-login" :class="{show : changeMode}" alt="" />
            <img src="@/assets/images/auth/reset-submit.png" class="image img-register" :class="{show : !changeMode}" alt="" />
          </div>
        </div>
      </div>
    </div>
  </main>
  <Dialog v-if="showDialog" :show="showDialog"> <ResetForm></ResetForm> </Dialog>
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

main.form-2-mode .forms-wrap {
  left: 55%;
}

main.form-2-mode .carousel {
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

.img-login {
  transform: translate(0, -50px);
}

.img-register {
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
    margin-top: 70px;
    border-radius: 2rem;
  }

  .inner-box {
    padding: 1rem;
  }

  .carousel {
    padding: 1.5rem 1rem;
    border-radius: 1.6rem;
  }

}
</style>