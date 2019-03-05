export function getRateLimit() {
  return {
    "name": "rate_limit",
    "enabled": true,
    "config": {
      "limit": "10-S",
      "policy": "local",
      'redis': {
        dsn:'',
        prefix:'',
      }
    }
  }
}


export function getRateLimitForm(formData) {
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
      title: 'limit - eg: 5-S, 10-M, 1000-H',
      key: 'config.limit',
      defaultValue: formData.config.limit,
      type: 'input',
      rule: {required: true, message: '', trigger: 'blur'},
      props: {
        placeholder: ''
      }
    },
    {
      title: 'policy',
      key: 'config.policy',
      defaultValue: formData.config.policy,
      type: 'radio',
      rule: {required: true, message: 'Name?', trigger: 'blur'},
      options: [
        {value: 'local', text: 'local'},
        {value: 'redis', text: 'redis'},
      ]
    },

    {
      title: 'redis.dsn',
      key: 'config.redis.dsn',
      defaultValue: formData.config.redis.dsn,
      type: 'input',
      rule: {required: true, message: '', trigger: 'blur'},
      props: {
        placeholder: ''
      }
    },

    {
      title: 'redis.prefix',
      key: 'config.redis.prefix',
      defaultValue: formData.config.redis.prefix,
      type: 'input',
      rule: {required: true, message: '', trigger: 'blur'},
      props: {
        placeholder: ''
      }
    },

  ]
}
