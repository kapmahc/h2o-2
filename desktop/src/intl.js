import 'moment/locale/zh-cn'
import 'moment/locale/zh-tw'

import Cookie from 'js-cookie'

import dataEn from 'react-intl/locale-data/en'
import dataZh from 'react-intl/locale-data/zh'


const KEY = 'locale'

export const setLocale = (lng) => {
  localStorage.setItem(KEY, lng)
  Cookie.set(KEY, lng, {expires: 365})
  window.location.reload()
}


export const detectLocale = () => {
  switch (localStorage.getItem(KEY)) {
    case 'zh-Hans':
      return {
        locale: 'zh-Hans',
        data: dataZh,
        moment: 'zh-cn',
      }
    case 'zh-Hant':
      return {
        locale: 'zh-Hant',
        data: dataZh,
        moment: 'zh-tw',
      }
    default:
      return {
        locale: 'en-US',
        data: dataEn,
        moment: 'en',
      }
  }
}
