import { userService } from "../../services"
import { context } from "../helpers/context"
import { storage } from "../helpers/storage"

export async function dispatchGetMe () {
  const user = userService.getMe(context.state.accessToken)
  storage.setUser(user)
}

export async function dispatchLoadLocalStorageUser () {
  const user = storage.getUser()
  context.commit("user", user)

  return user
}
