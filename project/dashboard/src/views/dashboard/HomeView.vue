<script setup lang="ts">

import DashboardLayout from '@/layouts/DashboardLayout.vue'
import { createConnectTransport } from '@connectrpc/connect-web'
import { createPromiseClient } from '@connectrpc/connect'
import { AuthenticationService } from '@/proto/gateway/authentication/v1/authentication_connect'
import { DockerService } from '@/proto/gateway/docker/v1/containers_connect'

const transport = createConnectTransport({
  baseUrl: "http://localhost:50051",
});

const client = createPromiseClient(DockerService, transport);

const response = await client.listContainers({});


</script>

<template>
  <dashboard-layout>

    Dashboard..
    
  </dashboard-layout>
</template>
