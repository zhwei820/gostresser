export function getSelectPluginForm(defaultValue) {
  return [
    {
      title: 'Plugin Name',
      key: 'pluginName',
      defaultValue: defaultValue,
      type: 'radio',
      rule: {required: true, message: 'Name?', trigger: 'blur'},
      options: [
        {value: 'body_limit', text: 'Body Limit'},
        {value: 'cb', text: 'Circuit Breaker'},
        {value: 'rate_limit', text: 'Rate Limiting'},
        {value: 'oauth2', text: 'oauth2'},
      ]
    },
  ]
}
