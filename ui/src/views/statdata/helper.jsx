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
        },
        {
            title: 'Average (ms)',
            key: 'avg_response_time',
        },
        {
            title: 'Min (ms)',
            key: 'min_response_time',
        },
        {
            title: 'Max (ms)',
            key: 'max_response_time',
        },
        {
            title: 'Content Size (bytes)',
            key: 'avg_content_length',
        },
        {
            title: '# reqs/sec',
            key: 'current_rps',
        },

    ]
}
