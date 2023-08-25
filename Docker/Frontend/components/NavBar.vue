<template>
    <nav id="title-center" class="navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="/">
      <!-- <a href="https://www.flaticon.es/iconos-gratis/cuidado-de-mascotas" title="cuidado-de-mascotas iconos">Cuidado-de-mascotas iconos creados por photo3idea_studio - Flaticon</a> -->
      <img src="cuidado-de-mascotas.png" >
    </a>

    <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
    </a>
  </div>

  <div id="navbarBasicExample" class="navbar-menu">
    <div class="navbar-start">
      <NuxtLink class="navbar-item" to="/" id="bar">
        Inicio
      </NuxtLink>

      <NuxtLink class="navbar-item" to="/perros/" id="bar">
        Perros
      </NuxtLink>

      <NuxtLink class="navbar-item" to="/nuevos" id="bar">
        Nuevos
      </NuxtLink>
    </div>

    <div class="navbar-end">
      <div class="navbar-item">
        <div class="buttons">
        </div>
      </div>
    </div>
  </div>
</nav>
</template>
<script>
export default{
    data(){
        return{
            
        }
    }
}
</script>

<style>
#bar{
  color: whitesmoke;
}
#title-center{
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: cadetblue;
  color: whitesmoke;
}

</style>