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
  font-size: 2rem;
  line-height: 2.5rem;
  font-weight: 400;
  letter-spacing: 0;
  @media screen and (min-width: 321px) {
    font-size: 2.25rem;
    line-height: 2.75rem;
    font-weight: 300;
  }
  @media screen and (min-width: 673px) {
    font-size: 2.625rem;
    line-height: 3.125rem;
    font-weight: 300;
  }
  @media screen and (min-width: 1057px) {
    font-size: 3rem;
    line-height: 3.5rem;
    font-weight: 300;
  }
  @media screen and (min-width: 1313px) {
    font-size: 3.75rem;
    line-height: 4.375rem;
    font-weight: 300;
  }
}
</style>