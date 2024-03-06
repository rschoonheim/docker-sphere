<script setup lang="ts">

const props = defineProps({
  size: {
    type: String,
    default: 'sm'
  }
})

function classes() {
  return {
    'sm': 'h-[32px]',
    'md': 'h-[38px]',
    'l': 'h-[40px]',
    'l2': 'h-[40px]',
    'xl': 'h-[64px]',
    'xl2': 'h-[80px]'
  }[props.size]
}

</script>

<template>
  <button
    role="button"
    class="text-[14px] w-fit pl-spacing-05 pr-[64px] bg-button-primary hover:bg-button-primary-hover active:bg-button-primary-active focus:bg-focus text-start"
    :class="classes()"
  >
    <slot />
  </button>
</template>

<style scoped>

</style>