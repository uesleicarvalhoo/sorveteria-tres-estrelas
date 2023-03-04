import { authService } from "../../services"
import { dispatchNotification, dispatchApiError } from "./notification"
import { context, storage } from "../helpers"
import { dispatchGetMe, dispatchLoadLocalStorageUser } from "./users"
import { router } from "../routers"

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

async function actionRefresh () {
  const response = await authService.refreshAcessToken(context.state.refreshToken)
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

export const dispatchLogin = async (email, password) => {
  try {
    const response = await authService.loginGetToken(email, password)
    const data = response.data

    if (data) {
      await actionLoggedIn(data)
    } else {
      dispatchNotification("Erro ao obter token", "Ocorreu um problema ao obter o seu Token de acesso, por favor entre em contato com o administrador do sistema", "danger")
    }

    await dispatchGetMe()
    router.push({ name: "home" })
  } catch (error) {
    await dispatchApiError(error)
    await dispatchLogout()
  }
}

export const dispatchRefreshToken = async () => {
  try {
    await actionRefresh()
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
