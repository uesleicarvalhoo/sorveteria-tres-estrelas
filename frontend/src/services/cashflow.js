import axios from "axios"
import { apiUrl } from "../config"
import { authHeaders } from "./auth"

export const cashFlowService = {
  async getBetween (start, end, token) {
    const headers = authHeaders(token)

    headers.params = {
      startAt: start,
      endAt: end
    }
    const res = await axios.get(`${apiUrl}/cashflow/`, headers)

    return res.data
  },
  async getAll (token) {
    const res = await axios.get(`${apiUrl}/cashflow/`, authHeaders(token))
    return res.data
  }
}
