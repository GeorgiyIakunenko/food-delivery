<script setup>
  import Button from "@/components/UI/Button.vue";
 const props = defineProps({
   supplier: {
     type: Object,
     required: true
   }
 })

  console.log(props.supplier)


</script>

<template>
  <div class="supplier-card">
    <div class="supplier-card__top">
      <p>{{supplier.type}}</p>
      <p>{{supplier.open_time}} - {{supplier.close_time}}</p>
    </div>
    <img src="@/assets/images/suppliers/aldi.jpg" alt="{{supplier.name}}">
    <p>{{ supplier.description }}</p>
    <div class="supplier-card__bottom">
      <h3>{{ supplier.name }}</h3>
      <Button intent="secondary">Open</Button>
    </div>

  </div>
</template>

<style scoped>
  .supplier-card {
    background-color: white;
    width: 300px;
    height: 400px;
    display: flex;
    font-family: "DM Sans", sans-serif;
    text-transform: capitalize;
    flex-direction: column;
    border-radius: 10px;
    padding: 10px;
    overflow: hidden;
    box-shadow: 0 0 10px rgba(0,0,0,0.2);
    transition: all 0.3s ease-in-out;
  }

  .supplier-card:hover {
    transform: scale(1.05);
  }

  .supplier-card img {
    width: 100%;
    height: 200px;
    object-fit: cover;
  }

  .supplier-card__bottom {
    margin-top: auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem;
    font-weight: bold;
    color : #fff;
    background-color: rgba(37, 114, 76, 0.37);
    border-radius: 10px;
  }

  p {
    font-size: 1rem;
    color: #151111;
  }

  .supplier-card__top {
    display: flex;
    justify-content: space-between;
    width: 100%;
    margin-bottom: 1rem;
  }
</style>