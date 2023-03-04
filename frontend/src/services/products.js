import { apiUrl } from "../config"
import { authHeaders } from "./auth"
import axios from "axios"

export const productsService = {
  async create (token, payload) {
    const res = await axios.post(`${apiUrl}/products/`, payload, authHeaders(token))

    return res.data
  },
  async getAll (token) {
    const res = await axios.get(`${apiUrl}/products/`, authHeaders(token))

    return res.data
  },
  async update (token, payload) {

    const res = await axios.post(`${apiUrl}/products/${payload.id}`, payload, authHeaders(token))
    return res.data
  },
  async delete (token, itemId) {
    const res = await axios.delete(`${apiUrl}/products/${itemId}`, authHeaders(token))

    return res.data
  }

}
