<template>
  <div class="wrapper">
    <div class="title">Sus vacunas</div>
    <div class="table">
      <thead>
        <tr>
          <th></th>
          <th>Nombre</th>
          <th>Fecha</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="vaccine in vaccines" :key="id_vaccine">
          <td>{{ owner.nombre }}</td>
          <td>{{ owner.fecha }}</td>
        </tr>
      </tbody>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      nombre: "",
      fecha: "",
      vaccines: [],
    };
  },
  methods: {
    async getVaccineByDogID(id) {
      const path = process.env.BACK_URL + "/dog/vaccines/" + id.toString();
      try {
        const res = await this.$axios.get(path);
        const data = res.data;
        this.vaccines = data;
      } catch (error) {
        console.log(error);
      }
    },
  },
  async created() {
    await this.getVaccinesByDogID(this.idDog);
  },
};
</script>
