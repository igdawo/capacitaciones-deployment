<template>
    <div class="box container"> 
        <div  id="box-dog-list">
            <table class="table is-hoverable is-fullwidth" id="table-box">
                <thead id="header-table">
                    <tr>
                        <th>Nombre</th>
                        <th>Raza</th>
                    </tr>
                </thead>
            <tbody id="table-body">
                <tr id="table-btn" v-for="(Dog, index) in dogs" :key="Dog._id" @click="fetchDogData(Dog._id)"> 
                    <td>{{Dog.Nombre}}</td>
                    <td>{{Dog.Raza}}</td>
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
        /*Recupera todos los perros de la BD*/ 
        async fetchAllDogs() {
            const path = process.env.BACK_URL + '/dog/';
            try {
            const res = await this.$axios.get(path);
            const data = res.data;
            this.dogs = data;
            }  
            catch (error) {
            console.log(error);
            }
        },

        /*Redirecciona a la ruta con los datos del perro utilizando su ID*/ 
        fetchDogData(dogID){
            this.$router.push('/perros/'+ dogID);
        }
    },
    async created(){
        await this.fetchAllDogs();
    }
}
</script>
<style>
#box-dog-list{
    height: 450px;
    overflow-y: scroll;
    
}

#table-btn:hover, #nav-btn:focus{
  color: white;
  background-color: #60b8b8;
  transition: 200ms all linear;
}

</style>