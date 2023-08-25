<template>
  <div class="card">
<header class="card-header" id="bar-vaccines">
  <p class="card-header-title" id="white-text">
      Vacunas
  </p>
  <button class="card-header-icon" aria-label="more options">
    <span class="icon">
      <i class="fas fa-angle-down" aria-hidden="true"></i>
    </span>
  </button>
</header>
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
                  :key="vaccine._id"> 
                  <td>{{vaccine.NombreVacuna}}</td>
                  <td>{{vaccine.Fecha }}</td>
              </tr>
          </tbody>
      </table>
  </div>
</div>

  <div class="modal"  :class="{'is-active': showModalFlag}">
<div class="modal-background"></div>
<div class="modal-card">
  <header class="modal-card-head" id="color-header">
    <p class="modal-card-title" id="white-text">Agregar una Vacuna</p>
    <button class="delete" aria-label="close" @click="cancelModal"></button>
  </header>
  <section class="modal-card-body">
      <div class="field">
          <label class="label">Nombre</label>
          <div class="control">
              <input class="input" type="text" placeholder="Nombre de la vacuna" v-model="vaccineName">
          </div>
      </div>
      <div class="field">
          <label class="label">Fecha</label>
          <div class="control">
              <input class="input" type="date" placeholder="yyyy-mm-dd" v-model="date">
          </div>
      </div>
  </section>
  <footer class="modal-card-foot">
    <button class="button is-primary" @click="okModal(), reloadVaccines()" >Guardar</button>
    <button class="button is-danger" @click="cancelModal" >Cancelar</button>
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
          vaccines:[]
      };
  },
  methods: {
      async fetchVaccinesByDogID(id) {
      const path = process.env.BACK_URL + '/dog/dog-vaccines/' + id.toString();
      try {
        const res = await this.$axios.get(path);
        const data = res.data;
        this.vaccines = data;
        }  
      catch (error) {
        console.log(error);
        }
    },

    async postVaccine() {
      let idDog = this.idDog;
      if(isNaN(idDog)){
        idDog = this.idDog
      }
      else{
        idDog = Number(this.idDog)
      }

      //Number(this.idDog)
      const date = new Date(this.date)
      const data = {
              NombreVacuna: this.vaccineName,
              Fecha: date,
              idPerro: idDog
          }
      const path = process.env.BACK_URL + '/vaccine/';
      try {
        const res = await this.$axios.post(path,data);
        }  
      catch (error) {
        console.log(error);
        }
    },

    reloadVaccines(){
      this.postVaccine();
      this.fetchVaccinesByDogID(this.idDog);
      this.date = ''
      this.vaccineName = ''
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
          this.date = ''
          this.vaccineName = ''
      }
  },
  async created(){
      await this.fetchVaccinesByDogID(this.idDog);
  }

}
</script>
<style>
#vaccine-list{
  height: 400px;
  overflow-y: scroll;
}
#color-header{
  background-color: #73CBCB;
}
#white-text{
  color: white;
}
#bar-vaccines{
  background-image: linear-gradient(to right, #eabf91 , #FCD3A8);
  
}
</style>