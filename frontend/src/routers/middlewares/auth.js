import { context } from "../../helpers/context"
import { dispatchCheckTokenExpiration } from "../../controller/auth"

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
    await dispatchCheckTokenExpiration()
    return next()
  }
}
