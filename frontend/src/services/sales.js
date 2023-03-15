import axios from "axios"
import { apiUrl } from "../config"
import { getContextHeaders } from "./utils"

export const salesService = {
  async registerSale (span, payload) {
    const res = await axios.post(`${apiUrl}/sales/`, payload, { headers: getContextHeaders(span) })

    return res.data
  },
  async getAll (span) {
    const res = await axios.get(`${apiUrl}/sales/`, { headers: getContextHeaders(span) })
    return res.data
  },
  async delete (span, payload) {
    const res = await axios.delete(`${apiUrl}/sales/${payload}`, { headers: getContextHeaders(span) })
    return res.data
  },
  async update (span, payload) {
    const res = await axios.patch(`${apiUrl}/sales/${payload.id}`, payload, { headers: getContextHeaders(span) })
    return res.data
  }
}
