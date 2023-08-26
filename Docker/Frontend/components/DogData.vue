<template>
  <wrapper class="div">
    <div class="title">Datos del perro</div>
    <div class="table">
      <thead>
        <tr>
          <th>Nombre</th>
          <th>Raza</th>
          <th>Color</th>
          <th>Edad</th>
        </tr>
      </thead>
      <tbody>
        <td>{{ dog.nombre }}</td>
        <td>{{ dog.raza }}</td>
        <td>{{ dog.color }}</td>
        <td>{{ dog.edad }}</td>
      </tbody>
    </div>
  </wrapper>
</template>

<script>
export default {
  data() {
    return {
      dog: {
        nombre: "",
        raza: "",
        color: "",
        edad: "",
      },
    };
  },
  methods: {
    async getDog(id) {
      const path = process.env.BACK_URL + "/dog/" + id.toString();
      try {
        const res = await this.$axios.get(path);
        const data = res.data;
        this.dog = dogResponde.data;
      } catch (error) {
        console.log(error);
      }
    },
  },
  async created() {
    await this.getDog(this.idDog);
  },
};
</script>
