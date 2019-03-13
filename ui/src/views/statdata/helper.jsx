// tslint:disable
export function cols(vm) {
    return [
        {
            title: 'Method',
            key: 'method',
        },
        {
            title: 'Name',
            key: 'name',
        },
        {
            title: '# requests',
            key: 'num_requests',
        },
        {
            title: '# fails',
            key: 'num_failures',
        },
        {
            title: 'Median (ms)',
            key: 'median_response_time',
            render: (h, params) => {
                return (
                    <div>{(params.row.median_response_time * 1000).toFixed(4)}</div>
                )
            },
        },
        {
            title: 'Average (ms)',
            key: 'avg_response_time',
            render: (h, params) => {
                return (
                    <div>{(params.row.avg_response_time * 1000).toFixed(4)}</div>
                )
            },

        },
        {
            title: 'Min (ms)',
            key: 'min_response_time',
            render: (h, params) => {
                return (
                    <div>{(params.row.min_response_time * 1000).toFixed(4)}</div>
                )
            },

        },
        {
            title: 'Max (ms)',
            key: 'max_response_time',
            render: (h, params) => {
                return (
                    <div>{(params.row.max_response_time * 1000).toFixed(4)}</div>
                )
            },

        },
        {
            title: 'Content Size (bytes)',
            key: 'avg_content_length',
        },
        {
            title: '# reqs/sec',
            key: 'rps',
        },

    ]
}
