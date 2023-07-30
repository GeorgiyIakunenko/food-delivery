<script setup>
  import { ref } from 'vue'
  const menuActive = ref(false)

  const toggleMenu = () => {
    menuActive.value = !menuActive.value
  }
  const CloseMenu = () => {
    menuActive.value = false
  }

</script>

<template>
  <header class="bg-primary-400 py-3 font-sans min-h-11 fixed top-0 left-0 right-0">
    <div class="container">
      <div class="wrapper flex justify-between items-center">
        <div class="bg-neutral-30 p-1 hover:bg-neutral-0 transition-all duration-300 cursor-pointer rounded-xl">
          <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 211 211" fill="none">
            <path d="M92.9631 0.625519C82.8044 0.625519 74.5404 8.88952 74.5404 19.0482V25.5985L86.8222 31.7394H92.9631H99.104L111.386 25.5985V19.0482C111.386 8.88952 103.122 0.625519 92.9631 0.625519Z" fill="#FF6A4D"/>
            <path d="M111.386 25.5985V19.0482C111.386 8.88952 103.122 0.625519 92.9631 0.625519V31.7394H99.104L111.386 25.5985Z" fill="#E64D2E"/>
            <path d="M0.440331 185.671C3.28602 199.685 16.2904 210.235 31.1448 210.235C45.9992 210.235 59.4129 199.685 62.2586 185.671L48.9173 179.53H8.96552L0.440331 185.671Z" fill="#474F54"/>
            <path d="M31.1448 197.953C20.9861 197.953 12.7221 189.689 12.7221 179.53C12.7221 169.371 20.9861 161.107 31.1448 161.107C41.3035 161.107 49.9768 169.371 49.9768 179.53C49.9768 189.689 41.3035 197.953 31.1448 197.953Z" fill="#F2F9FF"/>
            <path d="M62.2586 75.135H25.0039L12.7221 87.4168V124.262L25.0039 130.403H62.2586L80.6813 124.262V91.1013L74.5404 87.4168L62.2586 75.135Z" fill="#FF9100"/>
            <path d="M12.7221 62.8532V87.4168H74.5404L80.6813 84.2378L74.5404 81.2759V62.8532H12.7221Z" fill="#FABE2C"/>
            <path d="M99.104 185.671H0.440331V179.53C0.440331 155.789 20.0953 136.544 43.836 136.544L54.0708 130.403H111.386V175.436L99.104 185.671Z" fill="#FF6A4D"/>
            <path d="M0.440331 124.262H86.8222V136.544H0.440331V124.262Z" fill="#32393F"/>
            <path d="M129.808 155.335V178.424L123.668 185.671H99.104V136.544H86.8222C80.0672 136.544 74.5404 131.017 74.5404 124.262V93.5577H111.386V111.98C118.141 111.98 123.668 117.507 123.668 124.262V148.826L129.808 155.335Z" fill="#575F64"/>
            <path d="M129.808 155.335V178.424L123.668 185.671H99.104V136.544H92.9631V93.5577H111.386V111.98C118.141 111.98 123.668 117.507 123.668 124.262V148.826L129.808 155.335Z" fill="#474F54"/>
            <path d="M111.386 87.4168V81.2759C111.386 71.1172 103.122 62.8532 92.9631 62.8532C82.8044 62.8532 74.5404 71.1172 74.5404 81.2759V93.5577C74.5404 103.732 82.7884 111.98 92.9631 111.98H142.5C152.674 111.98 160.922 103.732 160.922 93.5577V87.4168H111.386Z" fill="#FF6A4D"/>
            <path d="M160.922 93.5577V87.4168H111.386V81.2759C111.386 71.1172 103.122 62.8532 92.9631 62.8532V111.98H142.5C152.674 111.98 160.922 103.732 160.922 93.5577Z" fill="#E64D2E"/>
            <path d="M99.104 31.7394H92.9631H86.8222L74.5404 37.8803V44.0212C74.5404 54.1798 82.8044 62.8532 92.9631 62.8532C103.122 62.8532 111.386 54.1798 111.386 44.0212V37.8803L99.104 31.7394Z" fill="#FFE1BA"/>
            <path d="M111.386 44.0212V37.8803L99.104 31.7394H92.9631V62.8532C103.122 62.8532 111.386 54.1798 111.386 44.0212Z" fill="#FFBFAB"/>
            <path d="M92.9631 25.5985H74.5404V37.8803H92.9631H123.668V25.5985H92.9631Z" fill="#474F54"/>
            <path d="M92.9631 25.5985H123.668V37.8803H92.9631V25.5985Z" fill="#32393F"/>
            <path d="M177.846 125.635C177.168 120.856 175.825 111.44 173.144 92.6877L148.64 81.2759V130.403C148.64 140.09 140.295 148.826 129.808 148.826H123.668V185.671H135.949H197.768L210.049 179.53V167.248C210.049 147.602 196.496 130.475 177.846 125.635Z" fill="#E64D2E"/>
            <path d="M185.486 99.6986H167.063C156.904 99.6986 148.64 91.4346 148.64 81.2759C148.64 71.1172 156.904 62.8532 167.063 62.8532H185.486V99.6986Z" fill="#FF6A4D"/>
            <path d="M181.19 149.012L179.405 155.836L169.026 150.738C157.182 155 148.64 166.237 148.64 179.53C148.64 196.46 162.415 210.235 179.345 210.235C196.275 210.235 210.049 196.46 210.049 179.53C210.049 163.233 197.246 149.988 181.19 149.012Z" fill="#32393F"/>
            <path d="M179.345 161.107C169.186 161.107 160.922 169.371 160.922 179.53C160.922 189.689 169.186 197.953 179.345 197.953C189.504 197.953 197.768 189.689 197.768 179.53C197.768 169.371 189.504 161.107 179.345 161.107Z" fill="#DFE7F4"/>
            <path d="M185.426 178.661L181.191 149.012C180.567 148.975 179.977 148.826 179.345 148.826C175.709 148.826 172.266 149.572 169.027 150.738L173.264 180.4L185.426 178.661Z" fill="#C7CFE1"/>
          </svg>
        </div>
        <nav class="menu max-[768px]:bg-neutral-200 rounded-e-xl shadow-2xl " :class='{ "active" : menuActive}'>
          <ul class="mt-12 md:mt-0 flex gap-3 text-neutral-0 md:gap-5 :gap-1">
            <li><router-link @click="CloseMenu" class="hover:text-red-600" to="/">Home</router-link></li>
            <li><router-link @click="CloseMenu" class="hover:text-red-600" to="/suppliers">Suppliers</router-link></li>
            <li><router-link @click="CloseMenu" class="hover:text-blue-600" to="/categories">Categories</router-link></li>
          </ul>
        </nav>
        <div class="left-block flex gap-5">
          <div class="user-box">
            <a class="hidden" href="#">Login</a>
            <a href="#">Register</a>
          </div>
          <div @click="toggleMenu" class="menu-burger md:hidden" :class='{ "active" : menuActive}'>
            <span class=""></span>
            <span class=""></span>
            <span class=""></span>
          </div>
        </div>

      </div>
    </div>
  </header>
</template>

<style scoped>

  header {
    z-index: 10;
    height: 70px;
  }

  /* Menu burger */

  .menu-burger {
    position: relative;
    width: 28px;
    height: 28px;
    padding: 0;
    background-color: transparent;
    border: none;
    cursor: pointer;
    outline: 0;
    transition: all 0.3s ease-in-out;
  }

  .menu-burger span {
    position: absolute;
    width: 100%;
    height: 2px;
    background-color: #fff;
    transition: all 0.3s ease-in-out;

  }

  .menu-burger span:nth-child(1) {
    top: 0;
  }

  .menu-burger span:nth-child(2) {
    top: 50%;
    transform: translateY(-50%);

  }

  .menu-burger span:nth-child(3) {
    bottom: 0;
  }

  .menu-burger.active span:nth-child(1) {
    top: calc(50% - 1px);
    transform: rotate(45deg);
  }

  .menu-burger.active span:nth-child(2) {
    opacity: 0;
  }

  .menu-burger.active span:nth-child(3) {
    bottom: calc(50% - 1px);
    transform: rotate(-45deg);
  }



  @media (max-width: 768px) {
    .menu {
      transform: translateX(-250%);
      transition: all 0.3s ease-in-out;
      left: 0;
      top: 0;
      bottom: 0;
      position: fixed;
    }

    .menu.active {
      transform: translateX(0);

    }

    .menu ul {
      flex-direction: column;
      justify-content: start;
      align-items: center;
      padding: 50px;
      height: 100%;
      font-size: 25px;
      grid-gap: 45px;
    }

  }



</style>