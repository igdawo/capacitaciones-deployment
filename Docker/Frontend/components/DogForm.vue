<template>
    <div class="box">
        <div class="columns is-justify-content-center">
            <div class="column is-5">
                <div class="field">
                    <label class="label text-color">Nombre</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="Ej. Perrito Bonito" v-model="dogName">
                    </div>
                </div>
                <div class="field">
                    <label class="label">Raza</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="Ej. Corg" v-model="dogBreed">
                    </div>
                </div>
                <div class="field">
                    <label class="label">Color</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="Ej. Blanco" v-model="dogColor">
                    </div>
                </div>
                <div class="field">
                    <label class="label">Edad</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="Ej. 15" v-model.number="dogAge">
                    </div>
                </div>
            
            </div>
            <div class="column is-5">
                <div class="field">
                    <label class="label text-color">Nombre del dueño</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="Ej. Juanito" v-model="ownerName">
                    </div>
                </div>
                <div class="field">
                    <label class="label">Edad del dueño</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="Ej. 20" v-model.number="ownerAge">
                    </div>
                </div>
                <div class="field">
                    <label class="label">Sexo del dueño</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="Ej. Masculino" v-model="ownerGender">
                    </div>
                </div>
                <div class="field is-grouped">
                <div class="control">
                    <button class="button is-primary" @click="sendData(), clearData()">Subir</button>
                </div>
                <div class="control">
                    <button class="button is-danger is-light" @click="clearData()">Cancelar</button>
                </div>
                </div>
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
            ownerName:'',
            ownerAge:'',
            ownerGender:''
        }
    },
    methods:{
        /*Envía los datos de un Perro a la BD*/ 
        async postDogData() {
             /*Crea la estructura de un perro*/ 
            const dogData = {
                Nombre: this.dogName,
                Edad: this.dogAge,
                Color: this.dogColor,
                Raza: this.dogBreed
            }
            const path =  process.env.BACK_URL + '/dog/';
            try {
            const res = await this.$axios.post(path,dogData);
            }  
            catch (error) {
            console.log(error);
            }
        },
         /*Envía los datos de un Dueño a la BD*/ 
        async postOwnerData() {

            const ownerData = {
                Nombre: this.ownerName,
                Edad: this.ownerAge,
                Sexo: this.ownerGender,
            }
            const path = process.env.BACK_URL + '/owner/';
                try {
                const res = await this.$axios.post(path,ownerData);
                }  
                catch (error) {
                console.log(error);
            }
        },
        /*Envía los datos del dueño y el perro creado a la BD*/ 
        async sendData(){
            this.postDogData();
            this.postOwnerData();
        },
        /*Limpia las variables utilizadas en el formulario*/ 
        clearData(){
            this.dogName='',
            this.dogAge='',
            this.dogColor='',
            this.dogBreed='',
            this.ownerName='',
            this.ownerAge='',
            this.ownerGender=''
        }
    }
}
</script>
<style>
.text-color{
    color: #64b3b3;
}
</style>