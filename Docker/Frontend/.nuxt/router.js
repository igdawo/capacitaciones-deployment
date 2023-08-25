import Vue from 'vue'
import Router from 'vue-router'
import { normalizeURL, decode } from 'ufo'
import { interopDefault } from './utils'
import scrollBehavior from './router.scrollBehavior.js'

const _dff72aa8 = () => interopDefault(import('..\\pages\\nuevos.vue' /* webpackChunkName: "pages/nuevos" */))
const _82e0258e = () => interopDefault(import('..\\pages\\perros\\index.vue' /* webpackChunkName: "pages/perros/index" */))
const _282b31f2 = () => interopDefault(import('..\\pages\\index.vue' /* webpackChunkName: "pages/index" */))
const _0579e9e1 = () => interopDefault(import('..\\pages\\perros\\_id.vue' /* webpackChunkName: "pages/perros/_id" */))

const emptyFn = () => {}

Vue.use(Router)

export const routerOptions = {
  mode: 'history',
  base: '/',
  linkActiveClass: 'nuxt-link-active',
  linkExactActiveClass: 'nuxt-link-exact-active',
  scrollBehavior,

  routes: [{
    path: "/nuevos",
    component: _dff72aa8,
    name: "nuevos"
  }, {
    path: "/perros",
    component: _82e0258e,
    name: "perros"
  }, {
    path: "/",
    component: _282b31f2,
    name: "index"
  }, {
    path: "/perros/:id",
    component: _0579e9e1,
    name: "perros-id"
  }],

  fallback: false
}

export function createRouter (ssrContext, config) {
  const base = (config._app && config._app.basePath) || routerOptions.base
  const router = new Router({ ...routerOptions, base  })

  // TODO: remove in Nuxt 3
  const originalPush = router.push
  router.push = function push (location, onComplete = emptyFn, onAbort) {
    return originalPush.call(this, location, onComplete, onAbort)
  }

  const resolve = router.resolve.bind(router)
  router.resolve = (to, current, append) => {
    if (typeof to === 'string') {
      to = normalizeURL(to)
    }
    return resolve(to, current, append)
  }

  return router
}
