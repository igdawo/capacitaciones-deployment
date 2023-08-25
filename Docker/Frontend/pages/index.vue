<template>
  <div class="container" id="homePage">
    <div class="columns " id="columnas">
  <div class="column  is-justify-content-center is-align-content-center">

  </div>
  <div id="title-center" class="column is-justify-content-center is-align-content-center box">
    <figure class="image is-128x128">
      <img class="is-rounded" src="../static/corgiicono.svg">
    </figure>
    <p class="is-size-3 pl-3"> Veterinaria Citiaps</p>
  </div>
  <div class="column is-justify-content-center is-align-content-center">

  </div>
</div>
  </div>
  
</template>

<script>
export default {
  name: 'IndexPage'
}
</script>
<style>
#homePage{
  height: 480px;
}
#title-center{
  background-color: #73CBCB;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 50%;
}
#columnas{
  background-color: white;
  height: 500px;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>