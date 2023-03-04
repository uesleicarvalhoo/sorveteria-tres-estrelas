import axios from "axios"
import { apiUrl } from "../config"
import { authHeaders } from "./auth"

export const salesService = {
  async registerSale (token, payload) {
    const res = await axios.post(`${apiUrl}/sales/`, payload, authHeaders(token))

    return res.data
  },
  async getAll (token) {
    const res = await axios.get(`${apiUrl}/sales/`, authHeaders(token))
    return res.data
  },
  async delete (token, payload) {
    const res = await axios.delete(`${apiUrl}/sales/${payload}`, authHeaders(token))
    return res.data
  },
  async update (token, payload) {
    const res = await axios.patch(`${apiUrl}/sales/${payload.id}`, payload, authHeaders(token))
    return res.data
  }
}
