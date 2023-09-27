import axios from "axios"
import { apiUrl } from "../config"
import { getContextHeaders } from "./utils"

export const transactionService = {
  async registerTransaction (span, payload) {
    const res = await axios.post(`${apiUrl}/transaction`, payload, { headers: getContextHeaders(span) })

    return res.data
  },

  async deleteTransaction (span, id) {
    const res = await axios.delete(`${apiUrl}/transaction/${id}`, { headers: getContextHeaders(span) })

    return res.data
  },

  async getBetween (span, start, end) {
    const res = await axios.get(`${apiUrl}/transaction/`, {
      headers: getContextHeaders(span),
      params: {
        startAt: start,
        endAt: end
      }
    })

    return res.data
  },
  async getAll (span) {
    const res = await axios.get(`${apiUrl}/transaction/`, { headers: getContextHeaders(span) })
    return res.data
  }
}
