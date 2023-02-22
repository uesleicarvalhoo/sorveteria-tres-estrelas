import axios from "axios"
import { apiUrl } from "../config"
import { authHeaders } from "./auth"

export const paymentsService = {
  async registerPayment (token, payload) {
    const res = await axios.post(`${apiUrl}/payments/`, payload, authHeaders(token))

    return res.data
  },

  async deletePayment (token, id) {
    const res = await axios.delete(`${apiUrl}/payments/${id}`, authHeaders(token))

    return res.data
  },

  async updatePayment (token, payload) {
    console.log(payload)
    const res = await axios.post(`${apiUrl}/payments/${payload.id}`, payload, authHeaders(token))
    return res.data
  },

  async getBetween (start, end, token) {
    const headers = authHeaders(token)

    headers.params = {
      startAt: start,
      endAt: end
    }
    const res = await axios.get(`${apiUrl}/payments/`, headers)

    return res.data
  },
  async getAll (token) {
    const res = await axios.get(`${apiUrl}/payments/`, authHeaders(token))
    return res.data
  }
}
