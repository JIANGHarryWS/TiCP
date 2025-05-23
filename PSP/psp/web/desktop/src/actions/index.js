/* Copyright (C) 2016-present, Yuansuan.cn */
import axios from 'axios'
import store from '../reducers'
import { dfApps } from '../utils'
import { gene_name } from '../utils/apps'
import { fetchDesktopAPP } from '@/domain/refreshDesktop'

import { env, appList, vis, currentUser } from '@/domain'
import { ListSessionRequest, SESSION_STATUS_ALL } from '@/domain/Vis'

import {
  allApps,
  generateCalcApp,
  GenerateDesktopApps,
  generateCloudApp
} from '@/utils'
export const dispatchAction = event => {
  const action = {
    type: event.target.dataset.action,
    payload: event.target.dataset.payload
  }

  if (action.type) {
    store.dispatch(action)
  }
}

export const refresh = async (pl, menu) => {
  if (menu.menus.desk[0].opts[4].check) {
    currentUser.fetch().then(async () => {
      await fetchDesktopAPP()
      store.dispatch({ type: 'DESKHIDE' })
    })

    setTimeout(() => {
      store.dispatch({ type: 'DESKSHOW' })
    }, 300)
  }
}

export const changeIconSize = (size, menu) => {
  let tmpMenu = { ...menu }
  tmpMenu.menus.desk[0].opts[0].dot = false
  tmpMenu.menus.desk[0].opts[1].dot = false
  tmpMenu.menus.desk[0].opts[2].dot = false
  let isize = 1

  if (size == 'large') {
    tmpMenu.menus.desk[0].opts[0].dot = true
    isize = 1.5
  } else if (size == 'medium') {
    tmpMenu.menus.desk[0].opts[1].dot = true
    isize = 1.2
  } else {
    tmpMenu.menus.desk[0].opts[2].dot = true
  }

  refresh('', tmpMenu)
  store.dispatch({ type: 'DESKSIZE', payload: isize })
  store.dispatch({ type: 'MENUCHNG', payload: tmpMenu })
}

export const deskHide = (payload, menu) => {
  let tmpMenu = { ...menu }
  tmpMenu.menus.desk[0].opts[4].check ^= 1

  store.dispatch({ type: 'DESKTOGG' })
  store.dispatch({ type: 'MENUCHNG', payload: tmpMenu })
}

export const changeSort = (sort, menu) => {
  let tmpMenu = { ...menu }
  tmpMenu.menus.desk[1].opts[0].dot = false
  tmpMenu.menus.desk[1].opts[1].dot = false
  tmpMenu.menus.desk[1].opts[2].dot = false
  if (sort == 'name') {
    tmpMenu.menus.desk[1].opts[0].dot = true
  } else if (sort == 'size') {
    tmpMenu.menus.desk[1].opts[1].dot = true
  } else {
    tmpMenu.menus.desk[1].opts[2].dot = true
  }

  refresh('', tmpMenu)
  store.dispatch({ type: 'DESKSORT', payload: sort })
  store.dispatch({ type: 'MENUCHNG', payload: tmpMenu })
}

export const changeTaskAlign = (align, menu) => {
  let tmpMenu = { ...menu }
  if (tmpMenu.menus.task[0].opts[align == 'left' ? 0 : 1].dot) return

  tmpMenu.menus.task[0].opts[0].dot = false
  tmpMenu.menus.task[0].opts[1].dot = false

  if (align == 'left') {
    tmpMenu.menus.task[0].opts[0].dot = true
  } else {
    tmpMenu.menus.task[0].opts[1].dot = true
  }

  store.dispatch({ type: 'TASKTOG' })
  store.dispatch({ type: 'MENUCHNG', payload: tmpMenu })
}

export const performApp = (act, menu) => {
  let data = {
    type: menu.dataset.action,
    payload: menu.dataset.payload
  }

  if (act == 'open') {
    if (data.type) store.dispatch(data)
  } else if (act == 'delshort') {
    if (data.type) {
      let apps = store.getState().apps
      let app = Object.keys(apps).filter(
        x =>
          apps[x].action == data.type ||
          (apps[x].payload == data.payload && apps[x].payload != null)
      )

      app = apps[app]
      if (app) {
        store.dispatch({ type: 'DESKREM', payload: app.name })
      }
    }
  }
}

export const delApp = (act, menu) => {
  let data = {
    type: menu.dataset.action,
    payload: menu.dataset.payload
  }

  if (act == 'delete') {
    if (data.type) {
      let apps = store.getState().apps
      let app = Object.keys(apps).filter(x => apps[x].action == data.type)
      if (app) {
        app = apps[app]
        if (app.pwa == true) {
          store.dispatch({ type: app.action, payload: 'close' })
          store.dispatch({ type: 'DELAPP', payload: app.icon })

          let installed = localStorage.getItem('installed')
          if (!installed) installed = '[]'

          installed = JSON.parse(installed)
          installed = installed.filter(x => x.icon != app.icon)
          localStorage.setItem('installed', JSON.stringify(installed))

          store.dispatch({ type: 'DESKREM', payload: app.name })
        }
      }
    }
  }
}

export const installApp = data => {
  let app = { ...data, type: 'app', pwa: true }

  let installed = localStorage.getItem('installed')
  if (!installed) installed = '[]'

  installed = JSON.parse(installed)
  installed.push(app)
  localStorage.setItem('installed', JSON.stringify(installed))

  let desk = localStorage.getItem('desktop')
  if (!desk) desk = dfApps.desktop
  else desk = JSON.parse(desk)

  desk.push(app.name)
  localStorage.setItem('desktop', JSON.stringify(desk))

  app.action = gene_name()
  store.dispatch({ type: 'ADDAPP', payload: app })
  store.dispatch({ type: 'DESKADD', payload: app })
  store.dispatch({ type: 'WNSTORE', payload: 'mnmz' })
}

export const getTreeValue = (obj, path) => {
  if (path == null) return false

  let tdir = { ...obj }
  path = path.split('.')
  for (let i = 0; i < path.length; i++) {
    tdir = tdir[path[i]]
  }

  return tdir
}

export const changeTheme = () => {
  var thm = store.getState().setting.person.theme,
    thm = thm == 'light' ? 'dark' : 'light'
  let icon = thm == 'light' ? 'sun' : 'moon'

  document.body.dataset.theme = thm
  store.dispatch({ type: 'STNGTHEME', payload: thm })
  store.dispatch({ type: 'PANETHEM', payload: icon })
  store.dispatch({ type: 'WALLSET', payload: thm == 'light' ? 0 : 1 })
}

export const loadSettings = () => {
  let sett = localStorage.getItem('setting') || '{}'
  sett = JSON.parse(sett)

  if (sett.person == null) {
    sett = JSON.parse(JSON.stringify(store.getState().setting))
    if (
      window.matchMedia &&
      window.matchMedia('(prefers-color-scheme: dark)').matches
    ) {
      sett.person.theme = 'dark'
    }
  }

  if (sett.person.theme != 'light') changeTheme()

  store.dispatch({ type: 'SETTLOAD', payload: sett })
}

// mostly file explorer
export const handleFileOpen = async (id, server) => {
  // handle double click open
  const item = store.getState().files.data.getId(id)
  let fetchFiles = null
  // 发送ls请求
  item?.type?.toLocaleLowerCase() === 'folder' &&
    (fetchFiles = await server.fetch(item.path))
  if (item != null) {
    if (item.type == 'folder') {
      store.dispatch({
        type: 'FILEDIR',
        payload: item.id,
        data: fetchFiles._children
      })
    }
  }
}
