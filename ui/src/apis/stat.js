import request from './request'

export const fetchStatData = (params = {}) => request.get('/stat/', {params})
