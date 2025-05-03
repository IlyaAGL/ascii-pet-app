<template>
  <form @submit.prevent="submit" class="form">
    <div class="form-group">
      <label>Name:</label>
      <input v-model="form.name" required />
    </div>
    <div class="form-group">
      <label>ASCII Art:</label>
      <textarea v-model="form.asciiArt" rows="6" required></textarea>
    </div>
    <div class="form-actions">
      <button type="button" @click="getPet">Get Pet</button>
      <button type="submit">Save Pet</button>
    </div>
  </form>
</template>

<script>
import { reactive } from 'vue'
import api from '../services/api'

export default {
  emits: ['saved'],
  setup(_, { emit }) {
    const form = reactive({ name: '', asciiArt: '' })

    const getPet = async () => {
      try {
        const { data } = await api.get('/pet')
        form.name = data.description
        form.asciiArt = data.ascii
      } catch (err) {
        console.error('Failed to fetch pet:', err)
      }
    }

    const submit = async () => {
    const payload = {
      ascii: form.asciiArt,
      description: form.name
    }
    await api.put('/pet', payload)
    emit('saved')
    form.name = ''
    form.asciiArt = ''
  }
    return { form, getPet, submit }
  }
}
</script>

<style scoped>
.form-group { margin-bottom: 15px }
label { display: block; font-weight: bold; margin-bottom: 5px }
input, textarea {
  width: 100%; padding: 8px; border-radius: 4px;
  border: 1px solid #ccc;
}
.form-actions {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
}
button {
  background: #4CAF50; color: #fff;
  border: none; padding: 10px 20px;
  border-radius: 4px; cursor: pointer;
}
button[type="button"] {
  background: #3498db;
}
button:hover { opacity: 0.9 }
</style>