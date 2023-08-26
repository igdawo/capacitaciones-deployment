<template>
  <nav
    class="navbar has-shadow is-link"
    role="navigation"
    aria-label="main navigation"
  >
    <!-- título de barra -->
    <div class="navbar-brand">
      <div class="navbar-item is-size-3"></div>
      <!-- formato hamburguesa -->
      <a
        role="button"
        class="navbar-burger"
        aria-label="menu"
        aria-expanded="false"
        data-target="navbarlinks"
      >
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
    </div>
    <!-- menú -->
    <div class="navbar-menu is-active" id="navbarlinks">
      <div class="navbar-end">
        <nuxt-link class="navbar-item" to="/inicio/">Inicio</nuxt-link>
        <nuxt-link class="navbar-item" to="/perros/">Perros</nuxt-link>
        <nuxt-link class="navbar-item" to="/nuevo/">Nuevo</nuxt-link>
      </div>
    </div>
  </nav>
</template>
