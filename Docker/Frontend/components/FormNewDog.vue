<template>
  <div class="container">
    <p class="title">Registrar un perro</p>
    <div class="field">
      <label class="label">Nombre del perro</label>
      <div class="control">
        <input
          class="input"
          type="text"
          placeholder="Nombre del perro"
          v-model="nombre"
        />
      </div>
    </div>
    <div class="field">
      <label class="label">Raza</label>
      <div class="control">
        <input
          class="input"
          type="text"
          placeholder="Raza del perro"
          v-model="raza"
        />
      </div>
    </div>
    <div class="field">
      <label class="label">Color</label>
      <div class="control">
        <input
          class="input"
          type="text"
          v-model="color"
          placeholder="Color del pelaje del perro"
        />
      </div>
    </div>
    <div class="field">
      <label class="label">Edad</label>
      <div class="control">
        <input
          class="input"
          type="text"
          placeholder="Edad del perro"
          v-model="edad"
        />
      </div>
    </div>
    <div class="field">
      <label class="label">Nombre del dueño</label>
      <div class="control">
        <input
          class="input"
          type="text"
          placeholder="Nombre del dueño del perro"
          v-model="dueno"
        />
      </div>
    </div>
    <div class="field">
      <div class="control">
        <label class="checkbox">
          <input type="checkbox" />
          Acepto los <a href="#">términos y condiciones</a>
        </label>
      </div>
    </div>

    <div class="field is-grouped">
      <div class="control">
        <button class="button is-link" @click="postDogData()">Registrar</button>
      </div>
      <div class="control">
        <button class="button is-link is-light" @click="clear()">
          Cancelar
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      nombre: "",
      raza: "",
      color: "",
      edad: "",
      dueno: "",
    };
  },
  methods: {
    async postDogData() {
      const dogData = {
        nombre: this.nombre,
        edad: this.edad,
        color: this.color,
      };
      const path = process.env.BACK_URL + "/dog";
      try {
        const res = await this.$axios.post(path, dogData);
      } catch (error) {
        console.log(error);
      }
    },
    clear() {
      (this.nombre = ""),
        (this.raza = ""),
        (this.color = ""),
        (this.edad = ""),
        (this.dueno = "");
    },
  },
};
</script>
