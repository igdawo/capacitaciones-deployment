<template>
    <div id="app">
        <NavBar/>
        <div class="container fluid" id="content-app" >
            <Nuxt/>
        </div>
        
        <Footer/>
    </div>

</template>
<script>
import NavBar from '@/components/NavBar.vue';
import Footer from '@/components/Footer.vue';
export default{
    components:{ NavBar,Footer}
}

</script>
<style>
html {
    caret-color: transparent;
    overflow-y: hidden;
}

#content-app{
    padding-top:1rem;
    padding-bottom: 1rem;
}
</style>