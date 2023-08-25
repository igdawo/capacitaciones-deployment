<template>
    <div class="box container"> 
        <div  id="box-dog-All">
            <table class="table is-hoverable is-fullwidth" id="table-box">
                <thead id="header-table">
                    <tr>
                        <th>Nombre</th>
                        <th>Raza</th>
                        <th>Edad</th>
                        <th>Color</th>
                        <th>Información</th>
                    </tr>
                </thead>
            <tbody id="table-body">
                <tr v-for="(Dog, index) in dogs" :key="Dog.id"> 
                    <td style="width: 20%">{{Dog.Nombre}}</td>
                    <td style="width: 20%">{{Dog.Raza}}</td>
                    <td style="width: 20%">{{Dog.Edad}}</td>
                    <td style="width: 20%">{{Dog.Color}}</td>
                    <td style="width: 20%; justify-content: center;align-items: center; "><button id="button" @click="fetchDogData(Dog.id)">>>></button></td>
                </tr>        
            </tbody>
        </table>
    </div>   
</div>
    
</template>
<script>
export default{
    data(){
        return{
            dogs:[]

        }
    }, 
    methods:{
        async fetchAllDogs() {
        const path = process.env.BACK_URL+'/perro/';
        try {
          const res = await this.$axios.get(path);
          const data = res.data;
          this.dogs = data;
          }  
        catch (error) {
          console.log(error);
          }
      },

        fetchDogData(dogID){
            this.$router.push('/perros/'+ dogID);
        }
    },
    created(){
        this.fetchAllDogs();
    }
}
</script>
<style>
#box-dog-All{
    height: 450px;
    overflow-y: scroll;
    
}
#button{
    width: 50%;
    background-color: cadetblue;
    color: whitesmoke;
    font-weight: bold;
    font-size: 100%;
}
</style>