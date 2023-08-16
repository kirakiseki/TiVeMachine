import type { LocationQueryRaw, Router } from 'vue-router'

import NProgress from 'nprogress' // progress bar

import { useUserStore } from '~/store'
import { isLogin } from '~/utils/auth'

export default function setupLoginGuard(router: Router) {
  router.beforeEach(async (to, from, next) => {
    NProgress.start()
    const userStore = useUserStore()
    if (isLogin()) {
      try {
        await userStore.info()
        next()
      }
      catch (error) {
        await userStore.logout()
        next({
          name: 'login',
          query: {
            redirect: to.name,
            ...to.query,
          } as LocationQueryRaw,
        })
      }
    }
    else {
      if (to.meta.requiresAuth === false) {
        next()
        return
      }
      if (to.name === 'login') {
        next()
        return
      }
      next({
        name: 'login',
        query: {
          redirect: to.name,
          ...to.query,
        } as LocationQueryRaw,
      })
    }
    NProgress.done()
  })
}
