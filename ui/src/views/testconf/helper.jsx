// tslint:disable
export function cols(vm) {
    return [
        {
            title: 'Test Conf Name',
            key: 'name',
        },

        {
            title: 'Operations',
            key: 'action',
            align: 'center',
            render: (h, params) => {
                return (
                    <div>
                        <i-button type='primary' onClick={vm.stress.bind(vm, params.row)}>Stress</i-button>
                        <i-button type='primary' onClick={vm.edit.bind(vm, params.row)}>Edit</i-button>
                        <i-button type='error' onClick={vm.remove.bind(vm, params.row)}>Delete</i-button>
                    </div>)
            },
        },

    ]
}
