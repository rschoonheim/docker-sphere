<script setup lang="ts">

</script>

<template>
  <h3 class="fluid-heading-04">
    <slot></slot>
  </h3>
</template>

<style scoped>

.fluid-heading-04 {
  font-family: 'IBM Plex Sans', sans-serif;
  font-size: 1.75rem;
  line-height: 2.25rem;
  font-weight: 400;
  letter-spacing: 0;
  @media screen and (min-width: 1057px) {
    font-size: 2rem;
    line-height: 2.5rem;
  }
}
</style>