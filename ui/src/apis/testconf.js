import request from './request'

export const fetchTestConf = (params = {}) => request.get('/conf', {params})
export const fetchTestConfByID = (id, params = {}) => request.get(`/conf/${id}/`, {params})
export const deleteTestConfByID = (id, params = {}) => request.delete(`/conf/${id}/`, {params})
export const putTestConf = (id, payload) => request.put(`/conf/${id}/`, payload)

export const createTestConf = (payload = {}) => request.post('/conf/', payload)
