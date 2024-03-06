<script setup lang="ts">
import HeadingThree from '@/components/headings/HeadingThree.vue'
import InputText from '@/components/input/InputText.vue'
</script>

<template>
  <div class="flex flex-row w-screen">
    <div class="bg-layer-01 h-screen min-w-fit w-full md:w-2/6 p-spacing-05 flex flex-col">

      <heading-three>
        Docker Sphere
      </heading-three>

      <input-text />

    </div>
    <div class="w-full p-spacing-05 hidden md:block">
    </div>
  </div>
</template>
