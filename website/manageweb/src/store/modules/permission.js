import { asyncRoutes, asyncRoutes2, constantRoutes } from '@/router'
// import { getMenus } from '@/api/user'
/* Layout */
import Layout from '@/layout'
// import { getToken } from '@/utils/auth'
/**
 * 通过meta.role判断是否与当前用户权限匹配
 * @param roles
 * @param route
 */
function hasPermission(roles, route) {
  if (route.meta && route.meta.roles) {
    return roles.some(role => route.meta.roles.includes(role))
  } else {
    return true
  }
}

/**
 * 递归过滤异步路由表，返回符合用户角色权限的路由表
 * @param routes asyncRoutes
 * @param roles
 */
export function filterAsyncRoutes(routes, roles) {
  const res = []

  routes.forEach(route => {
    const tmp = { ...route }
    if (hasPermission(roles, tmp)) {
      if (tmp.children) {
        tmp.children = filterAsyncRoutes(tmp.children, roles)
      }
      res.push(tmp)
    }
  })

  return res
}

// 在这里定义state 状态管理
const state = {
  token02: 'aabb02', // getToken(),
  routes: [],
  addRoutes: []
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = constantRoutes.concat(routes)
  }
}

const actions = {
  generateRoutes2({ commit }, roles) {
    return new Promise(resolve => {
      let accessedRoutes
      if (roles.includes('admin')) {
        // 如果是管理员则加载所有路由
        accessedRoutes = asyncRoutes
      } else {
        // 如果不是总管理员则进行过滤
        // accessedRoutes = filterAsyncRoutes(asyncRoutes, roles)
        accessedRoutes = asyncRoutes2
        return
        /**
        // 从服务端请求菜单列表数据
        // alert(state.token02)
        getMenus(state.token02).then(response => {
          const { code, menus } = response // 这种形式可以直接接收 {code:2000,data:"ok"}
          console.log(code)
          console.log(menus)
          // const { roles, name, avatar, introduction } = data
          // 处理路由、菜单
          // const asyncRouterMap_copy = []
          // accessedRoutes = asyncRouterMap_copy
        }).catch(error => {
          console.log(error)
          return
        })
        **/
        // end
      }
      commit('SET_ROUTES', accessedRoutes)
      resolve(accessedRoutes)
    })
  },
  generateRoutes({ commit }, data) {
    return new Promise(resolve => {
      // const { asyncRouterMap } = data
      // console.log('001')
      // console.log(data)
      const accessedRouters = convertRouter(data)
      // console.log('002')
      // console.log(accessedRouters)
      commit('SET_ROUTES', accessedRouters)
      resolve(accessedRouters)
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

/**
 *将后台的路由表进行格式化
 * @param {*} asyncRouterMap
 */
function convertRouter(asyncRouterMap) {
  const accessedRouters = []
  if (asyncRouterMap) {
    asyncRouterMap.forEach(item => {
      var isParent = false
      if (item.children) {
        isParent = true
      }
      var parent = generateRouter(item, isParent)
      // console.log('003')
      // console.log(parent)
      var children = []
      if (item.children) {
        // console.log('004')
        item.children.forEach(child => {
          // 下一级
          var children2 = []
          if (child.children) {
            child.children.forEach(child2 => {
              children2.push(generateRouter(child2, false))
            })
          }
          var parent2 = generateRouter(child, false)
          parent2.children = children2
          children.push(parent2)
          // 下一级 end
          // children.push(generateRouter(child, false))
        })
      }
      parent.children = children
      accessedRouters.push(parent)
    })
  }
  accessedRouters.push({ path: '*', redirect: '/404', hidden: true })
  return accessedRouters
}

// 对component的处理
function generateRouter(item, isParent) {
  var component = Layout // 多层嵌套时只能有一个Layout
  if (isParent !== true) {
    component = componentsMap[item.component]
  }
  var router = {
    path: item.path,
    name: item.name,
    meta: item.meta,
    // noCache: item.no_cache,
    hidden: item.hidden,
    // component: isParent ? Layout : () => import(item.component) // 这个不可以
    // component: isParent ? Layout : componentsMap[item.component]
    component: component
  }
  if (isParent !== true) {
    // router.meta = item.meta
    // router.name = item.name
  }
  // router.meta.noCache = true
  return router
}

// componentsMap 需要在事先定义好
export const componentsMap = {
  example_create: () => import('@/views/example/create'), // 添加文章
  example_edit: () => import('@/views/example/edit'), // 文章编辑
  table_index: () => import('@/views/tab/index'), // 表格首页
  Icon: () => import('@/views/svg-icons/index'), // 图标管理
  Menu: () => import('@/views/app/sys/menu'), // 菜单
  Admins: () => import('@/views/app/sys/admins'), // 后台管理员
  Role: () => import('@/views/app/sys/role'), // 后台角色
  Systemconfig: () => import('@/views/app/sys/systemconfig'), // 系统配置
  Emailconfig: () => import('@/views/app/sys/emailconfig'), // 邮箱设置配置
  Dbbackup: () => import('@/views/app/sys/dbbackup'), // 数据库备份
  Dbreduction: () => import('@/views/app/sys/dbreduction'), // 数据库还原
  Models: () => import('@/views/app/sys/models'), // 数据库还原
  Columns: () => import('@/views/app/sys/columns'), // 栏目列表
  Membersgroup: () => import('@/views/app/sys/membersgroup'), // 会员组
  Members: () => import('@/views/app/sys/members'), // 会员列表
  Adssite: () => import('@/views/app/sys/adssite'), // 广告位管理
  Ads: () => import('@/views/app/sys/ads'), // 广告管理
  Links: () => import('@/views/app/sys/links'), // 友情链接管理
  Messages: () => import('@/views/app/sys/messages'), // 留言管理
  Pizzlesclass: () => import('@/views/app/sys/pizzlesclass'), // 碎片分类
  Pizzles: () => import('@/views/app/sys/pizzles'), // 碎片分类
  Templates: () => import('@/views/app/sys/templates'), // 模板管理
  Weixinmenus: () => import('@/views/app/sys/weixinmenus'), // 微信目录
  Weixins: () => import('@/views/app/sys/weixins'), // 公众号管理
  Weixingroupmsg: () => import('@/views/app/sys/weixingroupmsg'), // 消息群发
  Weixinmsgmanager: () => import('@/views/app/sys/weixinmsgmanager'), // 消息管理
  Weixinfans: () => import('@/views/app/sys/weixinfans'), // 粉丝管理
  Weixinreplays: () => import('@/views/app/sys/weixinreplays'), // 自动回复
  Holidays: () => import('@/views/app/sys/holidays'), // 假日插件
  Adminlogs: () => import('@/views/app/sys/adminlogs'), // 管理员日志
  Test: () => import('@/views/app/sys/test') // test
}
