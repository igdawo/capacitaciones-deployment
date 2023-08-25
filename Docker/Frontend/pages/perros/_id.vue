<template>
    <div class="tile is-ancestor">
    
        <div class="tile is-child box">
            <p class="title">Datos del Perro</p>
            <p class="subtitle">Nombre: {{ dog.Nombre }}</p>
            <p>Raza: {{ dog.Raza }}</p>
            <p>Color: {{ dog.Color }}</p>
            <p>Edad: {{ dog.Edad }}</p>
        </div>
        <div class="tile is-child box">
            <p class="title">Datos del Dueño</p>
            <p class="subtitle">Nombre: {{ owner.Nombre }}</p>
            <p>Edad: {{ owner.Edad }}</p>
            <p>Sexo: {{ owner.Sexo }}</p>
            
        </div>
        <div class="tile is-child">
            <p class="title">Vacunas</p>
            <DogVaccines :idDog="$route.params.id"/>
        </div>
    
    </div>
</template>
<script>
import DogVaccines from '@/components/DogVaccines.vue';
import{mapGetters, mapActions,mapMutations} from 'vuex'


export default{
    data() {
        return {
            dog:{
                Nombre:'',
                Raza:'',
                Color:'',
                Edad:''
            },
            owner:{
                Nombre:'',
                Edad:''
            },
            vaccines:[]
        };
    }, 
    
    components: { DogVaccines },
    methods:{

        async getDog(id) {
            const dogID = id.toString();
            console.log(dogID);
            const path = process.env.BACK_URL+'/perro/'+ dogID;
            try {
            const res = await this.$axios.get(path);
            const data = res.data;
            this.dog = data;
            
            }  
            catch (error) {
            console.log(error);
            }
        },

      async getOwnerByDog(id) {
        const path = process.env.BACK_URL+'/perro/dueno/' + id.toString();
        try {
          const res = await this.$axios.get(path);
          const data = res.data;
          this.owner = data;
          }  
        catch (error) {
          console.log(error);
          }
      },

      
        loadData(){
            const dogID = this.$route.params.id
            this.getDog(dogID);
            this.getOwnerByDog(dogID);
            
        }
    },
    async created(){
        await this.loadData();
    }


}
</script>

