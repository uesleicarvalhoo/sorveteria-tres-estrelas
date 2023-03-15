import axios from "axios"
import { apiUrl } from "../config"
import { getContextHeaders } from "./utils"

export function authHeaders (token) {
  return {
    headers: {
      Authorization: `Bearer ${token}`
    }
  }
}

export const authService = {
  async getAcessToken (span, email, password) {
    const params = new URLSearchParams()
    params.append("email", email)
    params.append("password", password)

    return await axios.post(`${apiUrl}/auth/login`, params, { headers: getContextHeaders(span) })
  },
  async refreshAcessToken (span, token) {
    return await axios.post(`${apiUrl}/auth/refresh-token`, { refresh_token: token }, { headers: getContextHeaders(span) })
  }
}
