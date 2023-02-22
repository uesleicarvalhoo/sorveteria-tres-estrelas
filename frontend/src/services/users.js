import axios from "axios"
import { apiUrl } from "../config"
import { authHeaders } from "./auth"

export const userService = {
  async getMe (token) {
    const res = await axios.get(`${apiUrl}/users/me`, authHeaders(token))
    return res.data
  }
}
