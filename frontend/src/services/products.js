import axios from "axios"
import { apiUrl } from "../config"
import { getContextHeaders } from "./utils"

export const productsService = {
  async create (span, payload) {
    const res = await axios.post(`${apiUrl}/products/`, payload, { headers: getContextHeaders(span) })

    return res.data
  },
  async getAll (span) {
    const res = await axios.get(`${apiUrl}/products/`, { headers: getContextHeaders(span) })

    return res.data
  },
  async update (span, payload) {
    const res = await axios.post(`${apiUrl}/products/${payload.id}`, payload, { headers: getContextHeaders(span) })

    return res.data
  },
  async delete (span, itemId) {
    const res = await axios.delete(`${apiUrl}/products/${itemId}`, { headers: getContextHeaders(span) })

    return res.data
  }

}
