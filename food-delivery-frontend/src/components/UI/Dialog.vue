<script>
  export default {
    name : 'Dialog',
  }
</script>


<script setup>
  import Button from "@/components/UI/Button.vue";

  const props = defineProps({
    show : {
      type : Boolean,
      default : false
    },
    type : {
      type : String,
      validator(value) {
        return ['success', 'info', 'error'].includes(value)
      },
      default : 'info'
    }
  });

  const emit = defineEmits(['update:show'])

  const hideDialog = () => {
    emit('update:show', false)
  }
</script>

<template>
  <div class="dialog" @click.prevent="hideDialog" v-if="show">
    <div class="dialog__content" @click.stop>
      <div class="dialog__top">
        <div v-if="type === 'success'" class="dialog__top-type">
          <img class="type-img" src="@/assets/images/icons/success.svg" alt="success">
          <p>Success</p>
        </div>
        <div v-if="type === 'info'" class="dialog__top-type">
          <img class="type-img" src="@/assets/images/icons/info.svg" alt="info">
          <p>Info</p>
        </div>
        <div v-if="type === 'error'" class="dialog__top-type">
          <img class="type-img" src="@/assets/images/icons/error.svg" alt="error">
          <p>Error</p>
        </div>
        <button class="close-btn" @click="hideDialog">
          <svg width="25" height="25" viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M9.04958 7.98895L9.03853 8L9.04958 8.01105L12.9155 11.877C13.0069 11.9684 13.0069 12.1169 12.9155 12.2083L12.2083 12.9155C12.1169 13.0069 11.9684 13.0069 11.877 12.9155L11.216 12.2546L8.01105 9.04958L8 9.03853L7.98895 9.04958L4.1227 12.9155C4.0313 13.0069 3.88276 13.0069 3.79136 12.9155L3.08417 12.2083C2.99278 12.1169 2.99278 11.9684 3.08417 11.877L6.95042 8.01105L6.96147 8L6.95042 7.98895L3.08417 4.1227C2.99278 4.0313 2.99278 3.88276 3.08417 3.79136L3.79167 3.08418L3.79167 3.08417C3.88307 2.99278 4.03162 2.99278 4.12301 3.08417L7.98895 6.95042L8 6.96147L8.01105 6.95042L11.877 3.08449C11.9684 2.99309 12.1169 2.99309 12.2083 3.08449L12.9155 3.79167C13.0069 3.88307 13.0069 4.03162 12.9155 4.12301L12.2546 4.78395L9.04958 7.98895Z" fill="#9D9D99" stroke="#9D9D99" stroke-width="0.03125"/>
          </svg>
        </button>
      </div>
      <div class="dialog__slot"><slot></slot></div>
      <div v-if="type === 'success'" class="dialog__buttons">
        <Button @click="hideDialog" intent="secondary">Cancel</Button>
        <Button @click="hideDialog" intent="primary">Confirm</Button>
      </div>
      <div v-if="type === 'info'" class="dialog__buttons">
        <Button @click="hideDialog" intent="primary">Confirm</Button>
      </div>
      <div v-if="type === 'error'" class="dialog__buttons">
        <Button @click="hideDialog" intent="text">Cancel</Button>
        <Button @click="hideDialog" intent="danger">Ok</Button>
      </div>
    </div>
  </div>
</template>

<style scoped>
  .dialog {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    z-index: 120;
    width: 100%;
    height: 100%;
    background-color: rgba(0,0,0,0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 15px;
  }

  .close-btn {
    background-color: transparent;
    border: none;
    cursor: pointer;
    padding: 0;
    margin: 0;
    display: flex;
    transition: all 0.3s ease-in-out;
    align-items: center;
    outline: none;
  }

  .close-btn svg path {
    fill: #9D9D99;
    transition: all 0.3s ease-in-out;
  }

  .close-btn:hover svg path {
    fill: red;
    transform: translateY(1px);
    stroke: red;
  }

  .type-img {
    width: 25px;
    height: 25px;
    margin-right: 10px;
  }

  .dialog__content {
    display: flex;
    flex-direction: column;
    margin: auto;
    background-color: white;
    border-radius: 12px;
    min-height: 200px;
    min-width: 300px;
    //padding: 20px;
  }

  .dialog__slot  {
    padding: 20px;
    flex-grow: 1;
    max-width: 700px;
  }


  .dialog__buttons {
    margin-top: auto;
  }

  .dialog__top {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px;
  }

  .dialog__top-type {
    display: flex;
    align-items: center;
  }

  .dialog__buttons {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    grid-gap:  0 10px;
    background-color:  #F8F8F8;
    padding: 15px  20px 15px 15px;
    border-bottom-left-radius: 12px;
    border-bottom-right-radius: 12px;
  }

  img {
    width: 35px;
    height: 35px;
  }
</style>