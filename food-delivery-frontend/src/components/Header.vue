<script setup>
import {onMounted, ref} from 'vue'
import { RouterLink, RouterView } from 'vue-router'
import Button from "./Button.vue";
import logoIcon from "../assets/images/logo/logo.png";

const menuActive = ref(false)

const toggleMenu = () => {
  menuActive.value = !menuActive.value
}

onMounted(() => {
  console.log('mounted')
  console.log(logoIcon)
})


</script>

<template>
  <header class="header">
    <div class="container">
      <div class="header__wrapper">
        <router-link to="/"><img class="logo" src="../assets/images/logo/logo.png" alt="logo life upgrade" /></router-link>
        <nav class="menu" :class="{active : menuActive}">
          <ul class="menu-list">
            <li class="menu-list__item">
              <router-link class="menu-list__link" to="/">Home</router-link>
            </li>
            <li class="menu-list__item">
              <router-link class="menu-list__link" to="/suppliers">Suppliers</router-link>
            </li>
            <li class="menu-list__item">
              <a class="menu-list__link">Nutrition&WorkOut</a>
            </li>
            <li class="menu-list__item">
              <a class="menu-list__link" >MindSet</a>
            </li>
          </ul>
        </nav>
        <div class="button-box">
<!--          <button class="btn-contact btn-style">Login</button>-->
          <router-link to="/cart">
            <Button >Cart</Button>
          </router-link>
          <div @click="toggleMenu"  class="menu-btn" :class="{active : menuActive}">
            <span class="menu-btn__span menu-btn__span--1"></span>
            <span class="menu-btn__span menu-btn__span--2"></span>
            <span class="menu-btn__span menu-btn__span--3"></span>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>

<style scoped>
.header {
  position: fixed;
  top: 0;
  width: 100%;
  background-color: #fff;
  z-index: 999;
}

.header+.main {
  padding-top: 101px;
}

.header__wrapper {
  padding: 15px 0;
  justify-content: space-between;
  display: flex;
  align-items: center;
}

.menu-list {
  display: flex;

}

.menu-list__link {
  color: #000;
  text-decoration: none;
}

.menu-list__link.active {
  font-weight: 600;
  color: rgb(133, 18, 18);
  text-decoration: underline;

}

.menu-list__item+.menu-list__item {
  margin-left: 35px;
}

.btn-style {
  border: none;
  background: #25AB75;
  border-radius: 10px;
  font-weight: 600;
  color: #fff;
  padding: 10px 30px;
  cursor: pointer;
  box-shadow: 0 3px #999;
}

.btn-style:hover {
  background-color: #3e8e41
}

.btn-style:active {
  background-color: #3e8e41;
  box-shadow: 0 3px #666;
  transform: translateY(3px);
}

/* menu burger */

.button-box {
  display: flex;
  align-items: center;
}

.menu-btn {
  margin-left: 15px;
  z-index: 999;
  flex-direction: column;
  justify-content: space-between;
  height: 28px;
  width: 28px;
  cursor: pointer;
  display: none;
}



.menu-btn__span {
  height: 3px;
  width: 100%;
  background-color: #333;
  border-radius: 5px;
  transition: all 0.2s ease-in-out;
}


@media screen and (max-width: 768px) {

  .menu {
    display: block;
    position: fixed;
    top: 0;
    left: 0;
    height: 100vh;
    width: auto;
    padding: 15px;
    background-color: #f9f9f9;
    z-index: 998;
    transform: translateX(-110%);
    transition: transform 0.3s ease-in-out;
    box-shadow: 5px 0 10px 0 #aaa;
  }



  .menu-btn {
    display: flex;
  }

  .menu-list {
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin-top: 3rem;
  }

  .menu-list__item+.menu-list__item {
    margin-left: 0;
  }

  .menu-list__item {
    margin: 20px 0;
  }

  .menu-list__link {
    font-size: 1.5rem;
  }

  .menu-btn.active .menu-btn__span--1 {
    transform: rotate(45deg) translate(11px, 6px);
  }

  .menu-btn.active .menu-btn__span--2 {
    opacity: 0;
  }

  .menu-btn.active .menu-btn__span--3 {
    transform: rotate(-45deg) translate(12px, -7px);
  }

  .menu-btn__span {
    background-color: #333;
  }

  .menu.active {
    transform: translateX(0);
  }
}

@media (max-width: 500px) {
  .btn-contact {
    padding: 7px 15px;
  }

  .logo {
    max-width: 130px;
  }

  .menu-btn {
    margin-left: 10px;
  }

}

</style>