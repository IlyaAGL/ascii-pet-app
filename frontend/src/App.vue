<template>
  <div class="container">
    <h1>ASCII Pet App</h1>
    <PetForm @saved="fetchPet" @fetched="setPet" />
    <PetDisplay v-if="pet" :pet="pet" @deleted="pet = null" />
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import PetForm from './components/PetForm.vue'
import PetDisplay from './components/PetDisplay.vue'
import api from './services/api'

export default {
  components: { PetForm, PetDisplay },
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

    const setPet = (data) => {
      pet.value = data
    }

    onMounted(fetchPet)

    return { pet, fetchPet, setPet }
  }
}
</script>
