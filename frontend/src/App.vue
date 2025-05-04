<template>
  <div class="container">
    <h1>ASCII Pet App</h1>
    <PetForm @saved="fetchPet" />

    <transition name="fade">
      <PetDisplay
        v-if="pet"
        :pet="pet"
        @deleted="handleDeleted"
      />
    </transition>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import PetForm from './components/PetForm.vue'
import PetDisplay from './components/PetDisplay.vue'
import api from './services/api'

export default {
  name: 'App',
  components: { PetForm, PetDisplay },
  setup() {
    const pet = ref(null)

    const fetchPet = async () => {
      try {
        const { data } = await api.get('/pet')
        pet.value = data
      } catch (err) {
        console.error('Failed to fetch pet:', err)
        pet.value = null
      }
    }

    const handleDeleted = () => {
      pet.value = null
    }

    onMounted(fetchPet)
    return { pet, fetchPet, handleDeleted }
  }
}
</script>
