import axios from "axios"
import { apiUrl } from "../config"
import { authHeaders } from "./auth"

export const userService = {
  async getMe(token) {
    const res = await zipkinAxios.get(`${apiUrl}/auth/me`, authHeaders(token))
    return res.data
  }
}
