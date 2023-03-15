import axios from "axios"
import { apiUrl } from "../config"
import { getContextHeaders } from "./utils"

export const userService = {
  async getMe (span) {
    const res = await axios.get(`${apiUrl}/user/me`, { headers: getContextHeaders(span) })
    return res.data
  }
}
