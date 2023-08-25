<template>
  <div class="container" id="homePage">
    <div class="columns  ">
  <div class="column is-flex is-justify-content-center is-align-content-center">
    
  </div>
  <div id="title-center" class="column is-flex is-justify-content-center is-align-content-center box">
   <p id="title2">
    Veterinaria Citiaps
   </p>
  </div>
  <div class="column is-flex is-justify-content-center is-align-content-center">
    
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
html{overflow-y:hidden;}



#homePage{
  height: 480px;
}
#title-center{
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: cadetblue;
  color: whitesmoke;
  font-weight: bold;
}
#title2{
  font-size: 260%;
  font-family: "Open Sans", sans-serif;
  align-items: center;
  justify-content: center;
}
</style>