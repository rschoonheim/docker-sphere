<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import { createConnectTransport } from '@connectrpc/connect-web'
import { createPromiseClient } from '@connectrpc/connect'
import { DeploymentService } from '@/proto/api/deployment/deployment_connect'

const transport = createConnectTransport({
  baseUrl: "http://0.0.0.0:8080",
});

// Here we make the client itself, combining the service
// definition with the transport.
const client = createPromiseClient(DeploymentService, transport);

client.listDeployments({}).then((response) => {
  console.log(response);
});

</script>

<template>
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template>
