<script setup lang="ts">

import DashboardLayout from '@/layouts/DashboardLayout.vue'
import WidgetContainers from '@/components/widgets/WidgetContainers.vue'




</script>

<template>
  <dashboard-layout>

    <div class="grid grid-cols-2 lg:grid-cols-4">
      <widget-containers />
    </div>


  </dashboard-layout>
</template>
