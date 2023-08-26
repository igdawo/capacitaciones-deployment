<template>
  <div class="wrapper">
    <div class="title">Datos del dueño</div>
    <div class="table">
      <thead>
        <tr>
          <th>Nombre</th>
          <th>Edad</th>
          <th>Sexo</th>
        </tr>
      </thead>
      <tbody>
        <td>{{ owner.nombre }}</td>
        <td>{{ owner.edad }}</td>
        <td>{{ owner.sexo }}</td>
      </tbody>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      owner: {
        nombre: "",
        edad: "",
        sexo: "",
      },
    };
  },
  methods: {
    async getOwnerByDogID(id) {
      const path = process.env.BACK_URL + "/dog/owner/" + id.toString();
      try {
        const res = await this.$axios.get(path);
        const data = res.data;
        this.owner = data;
      } catch (error) {
        console.log(error);
      }
    },
  },
  async created() {
    await this.getOwnerByDogID(this.idDog);
  },
};
</script>
