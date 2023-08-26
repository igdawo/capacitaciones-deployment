<template>
  <div class="container">
    <div class="title">Perros registrados</div>
    <table class="table" id="tabledogs">
      <thead>
        <tr>
          <th>Nombre</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="dog in dogs" :key="dog.id">
          <td>
            {{ dog.nombre
            }}<button id="button" @click="fetchDogData(dog.id)"></button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "Perros",
  data() {
    return {
      dog: {
        nombre: "",
        raza: "",
        color: "",
        edad: "",
      },
      dogs: [],
    };
  },
  methods: {
    async fetchAllDogs() {
      const path = process.env.BACK_URL + "/perro/";
      try {
        const res = await this.$axios.get("path");
        const data = res.data;
        this.dogs = data;
      } catch (error) {
        console.log(error);
      }
    },
    fetchDogData(dogID) {
      this.$router.push("/perros/" + dogID);
    },
  },
};
</script>
