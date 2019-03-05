
// tslint:disable
export function cols(vm) {
  return [
    {
      title: 'Auth Name',
      key: 'name',
    },

    {
      title: 'Operations',
      key: 'action',
      width: 250,
      align: 'center',
      render: (h, params) => {
        return (
          <div>
            <i-button type='primary' onClick={vm.edit.bind(vm, params.row)}>Edit</i-button>
            <i-button type='error' onClick={vm.remove.bind(vm, params.row)}>Delete</i-button>
          </div>)
      },
    },

  ]
}
