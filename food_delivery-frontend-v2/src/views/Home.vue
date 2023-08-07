<script setup>
import { useUserStore } from "@/store/user";
import { onMounted, ref } from "vue";
import Filter from "@/components/Filter/Filter.vue";
import { getAllSuppliers } from "@/api/supplier";
import SupplierCard from "@/components/SupplierCard.vue";

const userStore = useUserStore();

const modalOpen = ref(false);
const modalType = ref("error");
const openModal = () => {
  modalOpen.value = true;
};
const closeModal = (value) => {
  modalOpen.value = false;
  alert(value);
};

const suppliers = ref([]);

onMounted(async () => {
  const res = await getAllSuppliers();
  console.log(res);
  if (res.success === true) {
    suppliers.value = res.data;
  }
});
</script>

<template>
  <main class="pt-24">
    <div class="">
      <div class="container">
        <Filter></Filter>
      </div>
    </div>

    <!--      Home

      <button @click="openModal">Open modal</button>
      <Modal
        :type="modalType"
        title="Failed to register"
        @modalClose="closeModal"
        v-bind:open="modalOpen"
        >Password is not correct</Modal
      >-->
    <div class="">
      <div class="container">
        <div
          class="mx-auto mt-7 grid max-w-fit grid-cols-2 gap-4 md:grid-cols-3 md:gap-7 xl:grid-cols-4"
        >
          <div v-for="supplier in suppliers" class="">
            <SupplierCard :supplier="supplier"></SupplierCard>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped></style>
