<script setup lang="ts">

</script>

<template>
  <h2 class="fluid-heading-03">
    <slot></slot>
  </h2>
</template>

<style scoped>
.fluid-heading-03 {
  font-family: 'IBM Plex Sans', sans-serif;
  font-size: 1.25rem;
  line-height: 1.75rem;
  font-weight: 400;
  letter-spacing: 0;
  @media screen and (min-width: 1313px) {
    font-size: 1.50rem;
  }
}
</style>