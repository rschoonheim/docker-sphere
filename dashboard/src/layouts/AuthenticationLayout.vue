<script setup lang="ts">

import InputText from '@/components/input/InputText.vue'
import HeadingThree from '@/components/headings/HeadingThree.vue'
</script>

<template>
  <div class="flex flex-row w-screen select-none overflow-hidden">
    <div class="bg-layer-01 h-screen min-w-fit w-full md:w-2/6 p-spacing-05 flex flex-col">

      <heading-three class="mb-spacing-05">
        Docker Sphere
      </heading-three>

      <slot />

    </div>
    <div class="w-full p-spacing-05 hidden md:block">
    </div>
  </div>
</template>

<style scoped>

</style>