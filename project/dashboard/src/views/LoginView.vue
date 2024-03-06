<script setup lang="ts">
import InputText from '@/components/input/InputText.vue'
import AuthenticationLayout from '@/layouts/AuthenticationLayout.vue'
import InputGroup from '@/components/input/InputGroup.vue'
import InputPassword from '@/components/input/InputPassword.vue'
import ButtonPrimary from '@/components/buttons/ButtonPrimary.vue'
import type { Ref } from 'vue'
import { ref } from 'vue'
import { required } from '@vuelidate/validators'
import useVuelidate from '@vuelidate/core'
import { createConnectTransport } from '@connectrpc/connect-web'
import { AuthenticationService } from '@/proto/gateway/authentication/v1/authentication_connect'
import { createPromiseClient } from '@connectrpc/connect'
import { useAuthenticationStore } from '@/stores/authentication'
import { type Router, useRouter } from 'vue-router'

const username: Ref<string> = ref<string>('')
const password: Ref<string> = ref<string>('')
const router: Router = useRouter()

const v$ = useVuelidate({
  username: {
    required: required
  },
  password: {
    required: required
  }
}, { username: username, password: password })


async function submit() {
  const validForm: boolean = await v$.value.$validate()
  if (!validForm) {
    return
  }

  const transport = createConnectTransport({
    baseUrl: "http://localhost:50051",
  });

  const client = createPromiseClient(AuthenticationService, transport);

  const response = await client.login({
    username: username.value,
    password: password.value,
  });

  useAuthenticationStore().setToken(response.token)

  router.push({
    name: 'dashboard'
  })

}
</script>

<template>
  <authentication-layout>

    <div class="flex flex-col gap-spacing-07">

      <input-group
        for="username"
        label="Username"
      >
        <input-text id="username" v-model="username" @keyup="v$.username.$validate()" />
      </input-group>

      <input-group
        for="password"
        label="Password"
      >
        <input-password id="password" v-model="password" @keyup="v$.password.$validate" />
      </input-group>

      <div class="flex flex-row justify-end">


        <button-primary size="md" class="w-1/2" @click="submit">
          Login
        </button-primary>
      </div>

    </div>

  </authentication-layout>
</template>
