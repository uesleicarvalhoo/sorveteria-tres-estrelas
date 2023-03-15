import axios from "axios"
import { apiUrl } from "../config"
import { getContextHeaders } from "./utils"

export const cashFlowService = {
  async getBetween (span, start, end) {
    const res = await axios.get(`${apiUrl}/cashflow/`, {
      params: {
        startAt: start,
        endAt: end
      },
      headers: getContextHeaders(span)
    })

    return res.data
  },
  async getAll (span) {
    const res = await axios.get(`${apiUrl}/cashflow/`, { headers: getContextHeaders(span) })
    return res.data
  }
}
