export function getBodylimit() {
  return           {
    "name": 'body_limit',
    "enabled": true,
    "config": {
      "limit": "40M"
    }
  }
}


export function getBodyLimitForm(formData) {
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
      title: 'Limit',
      key: 'config.limit',
      defaultValue: formData.config.limit,
      type: 'input',
      rule: {required: true, message: '1M, 100K ?', trigger: 'blur'},
      props: {
        placeholder: '1M, 100K ?'
      }
    },
  ]
}
