<script setup lang="ts">

</script>

<template>
  <input
    type="password"
    autocomplete="off"
    class="text-body-compact-01 px-spacing-05 text-text-primary text-[14px] bg-field-02 py-[13px]"
    @input="$emit('update:modelValue', $event.target.value)"
  />
</template>

<style scoped>

</style>