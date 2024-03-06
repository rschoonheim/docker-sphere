<script setup lang="ts">
import { createConnectTransport } from '@connectrpc/connect-web'
import { createPromiseClient } from '@connectrpc/connect'
import { DockerService } from '@/proto/gateway/docker/v1/containers_connect'
import { computed, type ComputedRef, onBeforeMount, ref } from 'vue'
import HeadingTwo from '@/components/headings/HeadingTwo.vue'

const transport = createConnectTransport({
  baseUrl: 'http://localhost:50051'
})

const client = createPromiseClient(DockerService, transport)

const containers = ref<{ id: string, name: string, status: string }[]>([])

onBeforeMount(() => {
  client.listContainers({}).then((response) => {
    containers.value = response.containers
  })
})

const onlineContainers: ComputedRef<number> = computed((): number => {
  return containers.value.filter((container) => {
    return container.status.startsWith('Up')
  }).length
})

// Count number of containers that exited
const exitedContainers: ComputedRef<number> = computed((): number => {
  return containers.value.filter((container) => {
    return container.status.startsWith('Exited')
  }).length
})

</script>

<template>
  <div class="bg-layer-01 hover:bg-layer-hover-01 transition transition-colors p-spacing-05 border border-border-tile-01 flex flex-col gap-spacing-05">
    <heading-two>Containers</heading-two>
    <p>Online: {{ onlineContainers }}</p>
    <p>Offline: {{ exitedContainers }}</p>
  </div>
</template>

<style scoped>

</style>