import { context } from "../../helpers/context"
import { dispatchRefreshToken } from "../../controller/auth"

export async function loginRequired (to, from, next) {
  if (to.name !== "login" && !context.state.loggedIn) {
    return next({ name: "login" })
  } else {
    return next()
  }
}

export async function refreshToken (to, from, next) {
  if (to.name === "login" || !context.state.loggedIn) {
    return next()
  } else {
    const now = Date.now() / 1000
    if ((now - 60) > context.state.expireTokenTime) {
      await dispatchRefreshToken()
    }

    return next()
  }
}
