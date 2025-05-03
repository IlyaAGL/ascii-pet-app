import axios from 'axios'

export default axios.create({
  baseURL: 'http://localhost:6060',
  headers: { 'Content-Type': 'application/json' }
})