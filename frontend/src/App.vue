<template>
    <div class="container">
      <h1>ASCII Pet App</h1>
      <PetForm @saved="fetchPet" />
    </div>
  </template>
  
  <script>
  import { ref, onMounted } from 'vue'
  import PetForm from './components/PetForm.vue'
  import api from './services/api'
  
  export default {
    components: { PetForm },
    setup() {
      const pet = ref(null)
      const fetchPet = async () => {
        try {
          const { data } = await api.get('/pet')
          pet.value = data
        } catch {
          pet.value = null
        }
      }
      onMounted(fetchPet)
      return { pet, fetchPet }
    }
  }
  </script>