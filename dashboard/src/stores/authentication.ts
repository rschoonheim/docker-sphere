import { ref, computed } from 'vue'
import type { Ref } from 'vue'
import { defineStore } from 'pinia'

export const useAuthenticationStore = defineStore('authentication', () => {

  const token: Ref<string> = ref<string | null>(null)


  return {
    token,
    setToken(newToken: string): void {
      token.value = newToken
    }
  }
})
