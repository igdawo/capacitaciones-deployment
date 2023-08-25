<template>
    <nav class="navbar" role="navigation" aria-label="main navigation" id="nav-back">
  <div class="navbar-brand">
    <a class="navbar-item" href="/" id="nav-btn">
      <figure class="image is-24x24">
        <img class="is-rounded" src="../static/corgiicono.svg">
      </figure>
    </a>

    <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
    </a>
  </div>

  <div id="navbarBasicExample" class="navbar-menu">
    <div class="navbar-start">
      <NuxtLink class="navbar-item" id="nav-btn" to="/">
        Inicio
      </NuxtLink>

      <NuxtLink class="navbar-item" id="nav-btn" to="/perros/">
        Perros
      </NuxtLink>

      <NuxtLink class="navbar-item" id="nav-btn" to="/nuevos">
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
#nav-back{
  background-image: linear-gradient(to right, #73CBCB , #6fe4e4);;
}
#nav-btn:hover, #nav-btn:focus{
  color: white;
  background-color: #529898;
  transition: 200ms all linear;
}
</style>