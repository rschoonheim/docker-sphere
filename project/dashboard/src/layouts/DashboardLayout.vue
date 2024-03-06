<script setup lang="ts"></script>

<template>
  <div class="flex flex-row w-screen select-none overflow-hidden bg-background">

    <div class="p-spacing-05 bg-layer-01 h-screen sticky top-0">
      Dashboard
    </div>

    <div class="w-full p-spacing-05">
      <slot />
    </div>


  </div>
</template>

<style scoped>

</style>