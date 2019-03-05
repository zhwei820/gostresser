export function getCB() {
  return {
    "name": "cb",
    "enabled": true,
    "config": {
      "name": "",
      "timeout": 1000,
      "max_concurrent_requests": 100,
      "error_percent_threshold": 50,
      "request_volume_threshold": 20,
      "sleep_window": 5000,
      "predicate": "statusCode == 0 || statusCode >= 500"
    }
  }
}


export function getCBForm(formData) {
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
      title: 'name',
      key: 'config.name',
      defaultValue: formData.config.name,
      type: 'input',
      rule: {required: true, message: '', trigger: 'blur'},
      props: {
        placeholder: ''
      }
    },
    {
      title: 'timeout',
      key: 'config.timeout',
      defaultValue: formData.config.timeout,
      type: 'input',
      rule: {required: true, message: '', trigger: 'blur'},
      props: {
        placeholder: ''
      }
    },
    {
      title: 'max_concurrent_requests',
      key: 'config.max_concurrent_requests',
      defaultValue: formData.config.max_concurrent_requests,
      type: 'input',
      rule: {required: true, message: '', trigger: 'blur'},
      props: {
        placeholder: ''
      }
    },
    {
      title: 'error_percent_threshold',
      key: 'config.error_percent_threshold',
      defaultValue: formData.config.error_percent_threshold,
      type: 'input',
      rule: {required: true, message: '', trigger: 'blur'},
      props: {
        placeholder: ''
      }
    },
    {
      title: 'request_volume_threshold',
      key: 'config.request_volume_threshold',
      defaultValue: formData.config.request_volume_threshold,
      type: 'input',
      rule: {required: true, message: '', trigger: 'blur'},
      props: {
        placeholder: ''
      }
    },
    {
      title: 'sleep_window',
      key: 'config.sleep_window',
      defaultValue: formData.config.sleep_window,
      type: 'input',
      rule: {required: true, message: '', trigger: 'blur'},
      props: {
        placeholder: ''
      }
    },
    {
      title: 'predicate',
      key: 'config.predicate',
      defaultValue: formData.config.predicate,
      type: 'input',
      rule: {required: true, message: '', trigger: 'blur'},
      props: {
        placeholder: ''
      }
    },
  ]
}
