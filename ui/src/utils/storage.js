import _get from 'lodash/get'
import _isEmpty from 'lodash/isEmpty'
import { TOKEN_NAME, USER_INFO } from '@/config'

const storage = window.localStorage || {
  setItem() {},
  getItem() {},
  removeItem() {},
  clear() {},
}

const KEYS = {
  tokenName: TOKEN_NAME,
  tokenInfo: USER_INFO,
}

function setItem(key, value) {
  storage.setItem(key, JSON.stringify(value || ''))
}

function getItem(key, paramPath, paramDefault = '') {
  const data = storage.getItem(key)
  if (data == null) return ''

  try {
    const value = JSON.parse(data)
    return _isEmpty(paramPath) ? value : _get(value, paramPath, paramDefault)
  } catch (e) {
    return data
  }
}

function clear() {
  storage.clear()
}

function removeItem(key) {
  storage.removeItem(key)
}

export default {
  KEYS,
  setItem,
  getItem,
  clear,
  removeItem,
}
