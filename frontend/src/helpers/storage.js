export const storage = {
  getAccessToken () {
    const token = localStorage.getItem("accessToken")
    return JSON.parse(token)
  },
  setAccessToken (token) {
    localStorage.setItem("accessToken", JSON.stringify(token))
  },
  removeAccessToken () {
    localStorage.removeItem("accessToken")
  },
  setUser (user) {
    localStorage.setItem("user", JSON.stringify(user))
  },
  removeUser () {
    localStorage.removeItem("user")
  },
  getUser () {
    const user = localStorage.getItem("user")
    return JSON.parse(user)
  }
}
