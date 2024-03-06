<script setup lang="ts">
defineProps({
  label: {
    type: String,
    required: true
  },
  for: {
    type: String,
    required: true
  }
})
</script>

<template>
  <div class="flex flex-col gap-spacing-02">
    <label class="text-[12px]" :for="$props.for">
      {{ label }}
    </label>
    <slot />
  </div>
</template>

<style scoped>

</style>