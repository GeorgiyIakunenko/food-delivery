<script setup>
  import {ref} from "vue";

  const props = defineProps({
    label: {
      type: String,
      required: true
    },
    type: {
      type: String,
      default: "text"
    },
    placeholder: {
      type: String,
      default: ""
    },
    name: {
      type: String,
      default: "",
      required: true
    },
  })

  // new feature from vue 3.3 version
  const modelValue = defineModel()

  const isPasswordHidden = ref(true)

</script>
<template>
  <div class="flex gap-2 flex-col items-start">
    <label :for="props.name" class="text-neutral-50 ml-3 font-sans font-medium capitalize">{{props.label}}</label>
    <div class="flex items-center justify-between bg-card-bg rounded-xl  px-3 border-0 w-full h-12">
      <input class="bg-transparent px-2 border-none outline-none w-full"
             :placeholder="props.placeholder"
             :label="props.label"
             v-model="modelValue"
             :id="props.name"
             :name="props.name"
             :type= "isPasswordHidden ?  props.type : 'text'"
             />
      <svg v-if="props.type ==='password'" class="cursor-pointer hover:-translate-y-0.5 transition-all " @click="isPasswordHidden = !isPasswordHidden" xmlns="http://www.w3.org/2000/svg" width="20" height="16" viewBox="0 0 20 16" fill="none">
        <path fill-rule="evenodd" clip-rule="evenodd" d="M10.0005 0C4.48046 0 0 4.84 0 8C0 11.086 4.57752 16 10.0005 16C15.4235 16 20 11.086 20 8C20 4.84 15.5195 0 10.0005 0ZM9.98349 11.984C7.77628 11.984 5.98129 10.19 5.98129 7.984C5.98129 5.778 7.77628 3.984 9.98349 3.984C12.1907 3.984 13.9857 5.778 13.9857 7.984C13.9857 10.19 12.1907 11.984 9.98349 11.984ZM12.0005 7.9996C12.0005 9.10417 11.1051 9.9996 10.0005 9.9996C8.89593 9.9996 8.0005 9.10417 8.0005 7.9996C8.0005 6.89503 8.89593 5.9996 10.0005 5.9996C11.1051 5.9996 12.0005 6.89503 12.0005 7.9996Z" fill="#42526E"/>
        <mask id="mask0_0_2695" style="mask-type:luminance" maskUnits="userSpaceOnUse" x="0" y="0" width="20" height="16">
          <path fill-rule="evenodd" clip-rule="evenodd" d="M10.0005 0C4.48046 0 0 4.84 0 8C0 11.086 4.57752 16 10.0005 16C15.4235 16 20 11.086 20 8C20 4.84 15.5195 0 10.0005 0ZM9.98349 11.984C7.77628 11.984 5.98129 10.19 5.98129 7.984C5.98129 5.778 7.77628 3.984 9.98349 3.984C12.1907 3.984 13.9857 5.778 13.9857 7.984C13.9857 10.19 12.1907 11.984 9.98349 11.984ZM12.0005 7.9996C12.0005 9.10417 11.1051 9.9996 10.0005 9.9996C8.89593 9.9996 8.0005 9.10417 8.0005 7.9996C8.0005 6.89503 8.89593 5.9996 10.0005 5.9996C11.1051 5.9996 12.0005 6.89503 12.0005 7.9996Z" fill="white"/>
        </mask>
        <g mask="url(#mask0_0_2695)">
        </g>
      </svg>
    </div>
    <div class="ml-3 text-red-400">
      <slot></slot>
    </div>
  </div>
</template>

<style scoped>

</style>