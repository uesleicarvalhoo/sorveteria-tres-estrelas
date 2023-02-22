import { context } from "../helpers/context"

export async function dispatchApiError(error) {
  let title = "Erro no servidor"
  let message = "Ocorreu um erro interno e não foi possível processar a sua solicitação!"
  const response = error.response

  if (response) {
    title = `${response.status} - ${title}`

    if (response.data.message) {
      message = response.data.message
    }
  } else {
    console.error(error)
  }

  await dispatchNotification(title, message, "danger")
}

export async function dispatchNotification(title, text, type = "info") {
  context.commit("showNotification", true)
  context.commit("notification", { title: title, text: text, type: type })
}

export async function dispatchConfirmNotification() {
  context.commit("showNotification", false)
}
