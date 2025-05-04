<template>
    <transition name="fade">
      <div v-if="pet" class="pet-container">
        <h2>{{ pet.description }}</h2>
        <textarea
          readonly
          rows="10"
          class="ascii-textarea"
          :value="pet.ascii"
        ></textarea>
        <button @click="handleDelete">Delete Pet</button>
      </div>
    </transition>
  </template>
  
  <script>
  import { defineComponent } from 'vue'
  import api from '../services/api'
  
  export default defineComponent({
    name: 'PetDisplay',
    props: {
      pet: {
        type: Object,
        required: true
      }
    },
    emits: ['deleted'],
    setup(_, { emit }) {
      const handleDelete = async () => {
        try {
          await api.delete('/pet')
          emit('deleted')
        } catch (err) {
          console.error('Failed to delete pet:', err)
        }
      }
      return { handleDelete }
    }
  })
  </script>
  
  <style scoped>
  .pet-container {
    text-align: center;
  }
  
  .ascii-textarea {
    width: 100%;
    font-family: 'Courier New', monospace;
    margin: 10px 0;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 4px;
    background: #fff;
  }
  
  button {
    margin-top: 10px;
    background: #e74c3c;
    color: #fff;
    border: none;
    padding: 8px 16px;
    border-radius: 4px;
    cursor: pointer;
  }
  button:hover {
    opacity: 0.9;
  }
  </style>
  