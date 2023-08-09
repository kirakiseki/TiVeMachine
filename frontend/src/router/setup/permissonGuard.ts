import type { RouteRecordNormalized, Router } from 'vue-router'
import NProgress from 'nprogress'

import usePermission from '~/hooks/permission'
import { useAppStore, useUserStore } from '~/store'
import { routes } from '~/main'

export default function setupPermissionGuard(router: Router) {
  router.beforeEach(async (to, from, next) => {
    const appStore = useAppStore()
    const userStore = useUserStore()
    const Permission = usePermission()
    const permissionsAllow = Permission.accessRouter(to)
    if (appStore.menuFromServer) {
      await appStore.fetchServerMenuConfig()
      const serverMenuConfig = [...appStore.appAsyncMenus]

      let exist = false
      while (serverMenuConfig.length && !exist) {
        const element = serverMenuConfig.shift()
        if (element?.name === to.name)
          exist = true

        if (element?.children) {
          serverMenuConfig.push(
            ...(element.children as unknown as RouteRecordNormalized[]),
          )
        }
      }
      if (exist && permissionsAllow)
        next()
      else next({ name: '404' })
    }
    else {
      if (permissionsAllow) {
        next()
      }
      else {
        const destination
                    = Permission.findFirstPermissionRoute(routes, userStore.role)
                    || { name: '404' }
        next(destination)
      }
    }
    NProgress.done()
  })
}
