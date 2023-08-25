<template>
    <div class="container is-flex is-justify-content-center" id="dogPage">
        <DogAll/>
    </div>
    
  </template>
  
  <script>
import DogAll from '@/components/DogAll.vue';

  export default {
    name: 'DogPage',
    data() {
        return {
        };
    },
    components: { DogAll }
}
  </script>
  <style>

  </style>