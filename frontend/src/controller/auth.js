import { authService } from "../services"
import { dispatchNotification, dispatchApiError } from "./notification"
import { context, storage } from "../helpers"
import { dispatchGetMe, dispatchLoadLocalStorageUser } from "./users"
import { router } from "../routers"
import { createSpan } from "../helpers/tracer"

async function actionLoggedIn (data) {
  context.commit("accessToken", data)
  context.commit("loggedIn", true)

  storage.setAccessToken(data)
}

async function actionLogout () {
  context.commit("accessToken", {})
  context.commit("loggedIn", false)
  storage.removeAccessToken()
}

async function actionRefresh (span) {
  const response = await authService.refreshAcessToken(span)
  const data = response.data

  if (data) {
    await actionLoggedIn(data)
  } else {
    await actionLogout()
  }
}

export const dispatchLogout = async () => {
  await actionLogout()
}

export const dispatchLogin = async (span, email, password) => {
  try {
    const response = await authService.getAcessToken(span, email, password)
    const data = response.data

    if (data) {
      await actionLoggedIn(data)
    } else {
      dispatchNotification("Erro ao obter token", "Ocorreu um problema ao obter o seu Token de acesso, por favor entre em contato com o administrador do sistema", "danger")
    }

    await dispatchGetMe(span)
    router.push({ name: "home" })
  } catch (error) {
    await dispatchApiError(error)
    await dispatchLogout()
  }
}

export const dispatchCheckTokenExpiration = async () => {
  const now = Date.now() / 1000
  if ((now - 60) > context.state.expireTokenTime) {
    await createSpan("refresh-access-token", async (span) => {
      await dispatchRefreshToken(span)
    })
  }
}

export const dispatchRefreshToken = async (span) => {
  try {
    await actionRefresh(span)
  } catch (error) {
    await dispatchApiError(error)
    await dispatchLogout()
  }
}

export const dispatchLoadContext = async () => {
  const token = await storage.getAccessToken()
  if (token) {
    await actionLoggedIn(token)
  }

  await dispatchLoadLocalStorageUser()
}
