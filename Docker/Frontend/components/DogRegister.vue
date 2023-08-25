<template>
    <div class="box">
        <div class="field">
            <label class="label">Nombre</label>
            <div class="control">
                <input class="input is-rounded" type="text" placeholder="Text input" v-model="dogName">
            </div>
        </div>
        <div class="field">
            <label class="label">Raza</label>
            <div class="control">
                <input class="input is-rounded" type="text" placeholder="Text input" v-model="dogBreed">
            </div>
        </div>
        <div class="field">
            <label class="label">Color</label>
            <div class="control">
                <input class="input is-rounded" type="text" placeholder="Text input" v-model="dogColor">
            </div>
        </div>
        <div class="field">
            <label class="label">Edad</label>
            <div class="control">
                <input class="input is-rounded" type="text" placeholder="Text input" v-model.number="dogAge">
            </div>
        </div>
        <div class="field">
            <label class="label">Selecciona un dueño:</label>
            <select v-model="selectedOwner">
                <option v-for="Owner in owners" :key="Owner.id" :value="Owner.id">
                    {{ Owner.Nombre }}
                </option>
            </select>
        </div>
    
    <div class="field is-grouped">
    <div class="control">
        <button id="button" class="button is-link" @click="postDogData(), clear()">Submit</button>
    </div>
    <div class="control">
        <button id="buttonInv" class="button is-link is-light" @click="clear()">Cancel</button>
    </div>
    </div>
</div>
</template>
<script>
export default{
    data(){
        return{
            dogName:'',
            dogAge:'',
            dogColor:'',
            dogBreed:'',
            owners: [],
            selectedOwner: null
        }
    },
    methods:{
        async postDogData() {

            const dogData = {
                Nombre: this.dogName,
                Edad: this.dogAge,
                Color: this.dogColor,
                Raza: this.dogBreed,
                Dueño_id: parseInt(this.selectedOwner)
            }
            const path = process.env.BACK_URL+'/perro/';
            try {
            const res = await this.$axios.post(path,dogData);
            }  
            catch (error) {
            console.log(error);
            }
        },
        async getOwners() {
            try {
                const path = process.env.BACK_URL+'/dueno/';
                const response = await this.$axios.get(path); 
                this.owners = response.data;
            } catch (error) {
                console.error('Error al obtener la lista de dueños:', error);
            }
        },
        clear(){
            this.dogName='',
            this.dogAge='',
            this.dogColor='',
            this.dogBreed='',
            this.selectedOwner=null
        }
    },
    async created(){
        await this.getOwners();
    },
}
</script>


<style>
#button{
    width: 100%;
    background-color: cadetblue;
    color: whitesmoke;
    font-weight: bold;
    font-size: 100%;
}
#buttonInv{
    width: 100%;
    background-color: whitesmoke;
    color: cadetblue;
    font-weight: bold;
    font-size: 100%;
}

</style>