import axios from "axios"
import { apiUrl } from "../config"
import { getContextHeaders } from "./utils"

export const paymentsService = {
  async registerPayment (span, payload) {
    const res = await axios.post(`${apiUrl}/payments/`, payload, { headers: getContextHeaders(span) })

    return res.data
  },

  async deletePayment (span, id) {
    const res = await axios.delete(`${apiUrl}/payments/${id}`, { headers: getContextHeaders(span) })

    return res.data
  },

  async updatePayment (span, payload) {
    const res = await axios.post(`${apiUrl}/payments/${payload.id}`, payload, { headers: getContextHeaders(span) })
    return res.data
  },

  async getBetween (span, start, end) {
    const res = await axios.get(`${apiUrl}/payments/`, {
      headers: getContextHeaders(span),
      params: {
        startAt: start,
        endAt: end
      }
    })

    return res.data
  },
  async getAll (span) {
    const res = await axios.get(`${apiUrl}/payments/`, { headers: getContextHeaders(span) })
    return res.data
  }
}
