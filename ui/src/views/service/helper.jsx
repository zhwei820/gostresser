
// tslint:disable
export function cols(vm) {
  return [
    {
      title: 'Api Name',
      key: 'name',
    },
    {
      title: 'Listen Path',
      key: 'listen_path',
      render: (h, params) => {
        return (
          <span>
            {params.row.proxy.listen_path}
          </span>)
      },
    },
    {
      title: 'Active',
      key: 'active',
      render: (h, params) => {
        const txt = '' + params.row.active
        const color = params.row.active ? 'success' : 'magenta'
        return (
          <span>
            <tag color={color}> {txt}</tag>
          </span>)
      },
    },
    {
      title: 'Operations',
      key: 'action',
      width: 250,
      align: 'center',
      render: (h, params) => {
        return (
          <div>
            <i-button type='primary' onClick={vm.copy.bind(vm, params.row)}>Copy</i-button>
            <i-button type='primary' onClick={vm.edit.bind(vm, params.row)}>Edit</i-button>
            <i-button type='error' onClick={vm.remove.bind(vm, params.row)}>Delete</i-button>
          </div>)
      },
    },

  ]
}
import {getBodyLimitForm, getBodylimit} from './plugins/bodylimit'
import {getCB, getCBForm} from './plugins/cb'
import {getRateLimit, getRateLimitForm} from './plugins/rate_limit'
import {getOauth, getOauthForm} from './plugins/oauth'

export function getFormList(obj, options=null) {
  return {
    body_limit : getBodyLimitForm,
    cb : getCBForm,
    rate_limit : getRateLimitForm,
    oauth2 : getOauthForm,
  }[obj.name].call(null, obj, options)
}

export function getDefaultFormData(name, ) {

  return {
    body_limit : getBodylimit,
    cb : getCB,
    rate_limit : getRateLimit,
    oauth2 : getOauth,
  }[name].call(null)
}
