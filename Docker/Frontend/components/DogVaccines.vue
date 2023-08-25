<template>
    <div class="card">
 
  <div class="card-content">
    <div class="content" id="vaccine-list">
        <table class="table is-hoverable is-fullwidth">
            <thead id="header-table">
                <tr>
                    <th>Nombre</th>
                    <th>Fecha</th>
                </tr>
            </thead>
            <tbody id="table-body">
                <tr v-for="(vaccine, index) in vaccines" 
                    :key="vaccine.id"> 
                    <td>{{vaccine.NombreVacuna}}</td>
                    <td>{{vaccine.Fecha }}</td>
                </tr>
            </tbody>
        </table>
    </div>
  </div>

  <div class="modal" :class="{'is-active': showModalFlag}">
  <div class="modal-background"></div>
  <div class="modal-card" id="popup">
    <header class="modal-card-head">
      <p class="modal-card-title">Agregar una Vacuna</p>
      <button class="delete" aria-label="close" @click="cancelModal"></button>
    </header>
    <section class="modal-card-body">
        <div class="field">
            <label class="label">Nombre</label>
            <div class="control">
                <input class="input is-rounded" type="text" placeholder="Text input" v-model="vaccineName">
            </div>
        </div>
        <div class="field">
            <label class="label">Fecha</label>
            <div class="control">
              <client-only><date-picker placeholder="MM/DD/YYYY" format="MM/dd/yyyy" v-model="date_today" /></client-only>
            </div>
        </div>
        
    </section>
    <footer class="modal-card-foot">
      <button class="button is-success" @click="okModal(), rechargeList()" >Save changes</button>
      <button class="button" @click="cancelModal" >Cancel</button>
    </footer>
  </div>
</div>

  <footer class="card-footer">
    <a href="#" class="card-footer-item js-modal-trigger" data-target="modal-js-example"  @click="showModal">Agregar una vacuna</a>
  </footer>
</div>
</template>
<script>

export default{
    props:['idDog'],
    data() {
        return {
            showModalFlag: false,
            okPressed: false,
            vaccineName: '',
            date:'',
            date_today:new Date(),
            vaccines:[]
        };
    },
    methods: {
        async getVaccinesByDog(id) {
        const path = process.env.BACK_URL+'/perro/vacunas/' + id.toString();
        try {
          const res = await this.$axios.get(path);
          const data = res.data;
          this.vaccines = data;
          }  
        catch (error) {
          console.log(error);
          }
      },

      async registerVaccine() {

        const date = new Date(this.date)
        const data = {
                NombreVacuna: this.vaccineName,
                Fecha: this.date_today,
                Perro_id: this.idDog
            }
        const path = process.env.BACK_URL+ '/vacuna/';
        try {
          const res = await this.$axios.post(path,data);
          }  
        catch (error) {
          console.log(error);
          }
      },

      rechargeList(){
        this.registerVaccine();
        this.getVaccinesByDog(this.idDog);
      },

        showModal() {
            this.okPressed = false;
            this.showModalFlag = true;
        },
        okModal() {
            this.okPressed = true;
            this.showModalFlag = false;
        },
        cancelModal() {
            this.okPressed = false;
            this.showModalFlag = false;
        }
    },
    async created(){
        await this.getVaccinesByDog(this.idDog);
    }

}
</script>
<style>
#vaccine-list{
    height: 400px;
    overflow-y: scroll;
}
#popup{
    height: 64%;
    overflow-y: scroll;
}
</style>