<template>
    <div class="container is-flex is-justify-content-center" id="dogPage">
        <DogList/>
    </div>
    
  </template>
  
  <script>
import DogList from '@/components/DogList.vue';

  export default {
    name: 'DogPage',
    data() {
        return {
        };
    },
    components: { DogList }
}
  </script>
  <style>

  </style>