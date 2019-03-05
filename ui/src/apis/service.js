import request from './request'

export const fetchService = (params = {}) => request.get('/apis', { params })
export const fetchServiceByName = (name, params = {}) => request.get(`/apis/${name}/`, { params })
export const deleteServiceByName = (name, params = {}) => request.delete(`/apis/${name}/`, { params })
export const putService = (name, payload) => request.put(`/apis/${name}/`, payload)

export const createService = (payload = {}) => request.post('/apis/', payload)
