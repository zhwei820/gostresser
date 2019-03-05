export function getOauth() {
  return           {
    "name": 'oauth2',
    "enabled": true,
    "config": {
      "server_name": ""
    }
  }
}


export function getOauthForm(formData, server_names) {
  return [
    {
      title: 'Name',
      key: 'name',
      defaultValue: formData['name'],
      type: 'html',
      rule: {required: true, message: 'Name?', trigger: 'blur'},
    },
    {
      title: 'Enabled',
      key: 'enabled',
      defaultValue: formData['enabled'],
      type: 'switch',
      rule: {required: true, message: 'Enabled?', trigger: 'blur'},
      props: {
        placeholder: 'Enabled?'
      }
    },
    {
      title: 'server_name',
      key: 'config.server_name',
      defaultValue: formData.config.server_name,
      type: 'select',
      options: server_names,
    },
  ]
}
