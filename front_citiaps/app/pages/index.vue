<script lang="ts" setup>

const { data: movies, error, status} = await useFetch('api/movie', {
    lazy: true,
});

</script>

<template>

</template>