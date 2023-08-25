<template>
    <div class="tile is-ancestor">
    <div class="tile is-4 is-vertical is-parent">
        <div class="tile is-child box">
            <p class="title is-size-4 text-color">Datos del Perro</p>
            <p class="subtitle text-color2">Nombre: {{ data.Nombre }}</p>
            <p>Raza: {{ data.Raza }}</p>
            <p>Color: {{ data.Color }}</p>
            <p>Edad: {{ data.Edad }}</p>
        </div>
        <div class="tile is-child box">
        <p class="title is-size-4 text-color">Datos del Dueño</p>
        <p class="subtitle text-color2">Nombre: {{ ownerData.Nombre }}</p>
        <p>Edad: {{ ownerData.Edad }}</p>
        <p>Sexo: {{ ownerData.Sexo }}</p>
            
        </div>
    </div>
    <div class="tile is-parent">
        <div class="tile is-child">
            <VaccineList :idDog="$route.params.id"/>
        </div>
    </div>
    </div>
</template>
<script>
import VaccineList from '@/components/VaccineList.vue';
import{mapGetters, mapActions,mapMutations} from 'vuex'


export default{
    data() {
        return {
            data :{
                Nombre:'',
                Raza:'',
                Color:'',
                Edad:''
            },
            ownerData:{
                Nombre:'',
                Edad:''
            },
            vaccines:[]
        };
    }, 
    components: { VaccineList },
    methods:{
         /*Recupera los datos de un perro a partir de su ID*/ 
        async fetchDogByID(id) {
        const dogID = id.toString();
        const path = process.env.BACK_URL + '/dog/'+ dogID ;
        try {
          const res = await this.$axios.get(path);
          const data = res.data;
          this.data = data;
          
          }  
        catch (error) {
          console.log(error);
          }
      },
      /*Recupera los datos deL dueño de un perro a partir de la ID del perro*/ 
      async fetchOwnerByDogID(id) {
        const path = process.env.BACK_URL + '/dog/dog-owner/' + id.toString();
        try {
          const res = await this.$axios.get(path);
          const data = res.data;
          this.ownerData = data;
          }  
        catch (error) {
          console.log(error);
          }
      },
      /*Carga la información a partir de la ID del perro*/ 
        loadData(){
            const dogID = this.$route.params.id
            this.fetchDogByID(dogID);
            this.fetchOwnerByDogID(dogID);
            
        }
    },
    async created(){
        await this.loadData();
    }


}
</script>
<style>
.text-color2{
    color: #6B3C35;
}
</style>