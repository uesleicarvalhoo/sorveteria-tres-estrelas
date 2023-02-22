import axios from "axios"
import { apiUrl } from "../config"

export function authHeaders (token) {
  return {
    headers: {
      Authorization: `Bearer ${token}`
    }
  }
}

export const authService = {
  async loginGetToken (email, password) {
    const params = new URLSearchParams()
    params.append("email", email)
    params.append("password", password)
    return await axios.post(`${apiUrl}/auth/login`, params)
  },
  async refreshAcessToken (token) {
    return await axios.post(`${apiUrl}/auth/refresh-token`, { refresh_token: token }, authHeaders(token))
  }
}
